package room

import (
	"fmt"
	"log/slog"
)

type Player struct {
	sign byte
}

type Room struct {
	clients map[*Client]Player

	register   chan *Client
	unregister chan *Client

	broadcast chan []byte

	id string
}

func NewRoom(id string) *Room {
	return &Room{
		clients: make(map[*Client]Player),

		register:   make(chan *Client),
		unregister: make(chan *Client),

		broadcast: make(chan []byte),

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

func (r *Room) registerClient(client *Client, store *RoomStore) bool {
	if len(r.clients) >= 2 {
		close(client.sendCh)
		return false
	}
	r.clients[client] = Player{
		sign: ' ',
	}

	store.roomsUpdateChan <- nil
	return len(r.clients) == 2
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
		client.sendCh <- []byte(fmt.Sprintf(`{"type": "start", "you": "%c", "turn": "%c"}`, lastSign, game.turn))
	}
}

func (r *Room) stopGame() {
	r.broadcastMsg([]byte(`{"type": "stop"}`))
}

func (r *Room) Run(store *RoomStore) {
	defer func() {
		store.roomUnregister <- r.id
	}()

	slog.Info("Room started", "id", r.id)

	game := NewGame()
	// TODO: Handle player leave
	for {
		select {
		case client, ok := <-r.register:
			if !ok {
				return
			}
			full := r.registerClient(client, store)
			if !full {
				continue
			}
			r.startGame(game)
		case client := <-r.unregister:
			r.unregisterClient(client, store)
			if len(r.clients) == 0 {
				return
			} else {
				r.stopGame()
			}
		case message := <-r.broadcast:
			r.broadcastMsg(message)
		}
	}
}
