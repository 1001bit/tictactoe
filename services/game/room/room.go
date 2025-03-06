package room

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"
)

type Player struct {
	sign byte
}

type ClientMsg struct {
	msg    []byte
	client *Client
}

type Room struct {
	clients map[*Client]Player

	register   chan *Client
	unregister chan *Client

	gameMsgCh chan ClientMsg

	id string
}

func NewRoom(id string) *Room {
	return &Room{
		clients: make(map[*Client]Player),

		register:   make(chan *Client),
		unregister: make(chan *Client),

		gameMsgCh: make(chan ClientMsg),

		id: id,
	}
}

func (r *Room) broadcastMsg(msg []byte) {
	for client := range r.clients {
		select {
		case client.sendCh <- msg:
		default:
			delete(r.clients, client)
			close(client.sendCh)
		}
	}
}

func (r *Room) registerClient(client *Client, store *RoomStore, game *Game) {
	if len(r.clients) >= 2 {
		close(client.sendCh)
		return
	}
	r.clients[client] = Player{
		sign: ' ',
	}

	store.roomsUpdateChan <- nil
	if len(r.clients) == 2 {
		r.startGame(game)
	}
}

func (r *Room) unregisterClient(client *Client, store *RoomStore) {
	if _, ok := r.clients[client]; !ok {
		return
	}
	delete(r.clients, client)
	close(client.sendCh)
	store.roomsUpdateChan <- nil
}

func (r *Room) startGame(game *Game) {
	game.Start()

	lastSign := byte(' ')
	for client := range r.clients {
		if lastSign == ' ' {
			lastSign = 'X'
		} else {
			lastSign = 'O'
		}

		r.clients[client] = Player{
			sign: lastSign,
		}
		client.sendCh <- fmt.Appendf([]byte{}, `{"type": "start", "you": "%c", "turn": "%c"}`, lastSign, game.turn)
	}
}

func (r *Room) broadcastStop() {
	r.broadcastMsg([]byte(`{"type": "stop"}`))
}

func (r *Room) broadcastEnd(result byte) {
	r.broadcastMsg(fmt.Appendf([]byte{}, `{"type": "end", "result": "%c"}`, result))
}

func (r *Room) handleGameMsg(msg ClientMsg, game *Game) {
	if game.turn != r.clients[msg.client].sign {
		return
	}

	type moveMsg struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	mm := moveMsg{}
	err := json.Unmarshal(msg.msg, &mm)
	if err != nil {
		return
	}

	ok := game.Place(mm.X, mm.Y)
	if !ok {
		return
	}

	res := game.CheckVictory(mm.X, mm.Y)
	r.broadcastMoveMsg(mm.X, mm.Y, r.clients[msg.client].sign, res)
	if res != ' ' {
		go func() {
			time.Sleep(3 * time.Second)
			r.startGame(game)
		}()
	}
}

func (r *Room) broadcastMoveMsg(x, y int, sign, result byte) {
	r.broadcastMsg(fmt.Appendf([]byte{}, `{"type": "move", "x": %d, "y": %d, "sign": "%c", "result": "%c"}`, x, y, sign, result))
}

func (r *Room) Run(store *RoomStore) {
	defer func() {
		store.roomUnregister <- r.id
	}()

	slog.Info("Room started", "id", r.id)

	game := NewGame()
	for {
		select {
		case client, ok := <-r.register:
			if !ok {
				return
			}
			r.registerClient(client, store, game)
		case client := <-r.unregister:
			r.unregisterClient(client, store)
			if len(r.clients) == 0 {
				return
			} else {
				r.broadcastStop()
			}
		case msg := <-r.gameMsgCh:
			r.handleGameMsg(msg, game)
		}
	}
}
