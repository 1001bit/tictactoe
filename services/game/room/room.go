package room

import "log/slog"

type Room struct {
	clients map[*Client]bool

	register   chan *Client
	unregister chan *Client

	broadcast chan []byte

	id string
}

func NewRoom(id string) *Room {
	return &Room{
		clients: make(map[*Client]bool),

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

func (r *Room) registerClient(client *Client, store *RoomStore) {
	if len(r.clients) >= 2 {
		close(client.sendCh)
		return
	}
	r.clients[client] = true

	store.roomsUpdateChan <- nil
}

func (r *Room) unregisterClient(client *Client, store *RoomStore) {
	if _, ok := r.clients[client]; !ok {
		return
	}
	delete(r.clients, client)
	close(client.sendCh)

	if len(r.clients) == 0 {
		return
	} else {
		r.broadcastMsg([]byte(`{"type": "playerLeft"}`))
		store.roomsUpdateChan <- nil
	}
}

func (r *Room) Run(store *RoomStore) {
	defer func() {
		store.roomUnregister <- r.id
	}()

	slog.Info("Room started", "id", r.id)
	for {
		select {
		case client, ok := <-r.register:
			if !ok {
				return
			}
			r.registerClient(client, store)
		case client := <-r.unregister:
			r.unregisterClient(client, store)
		case message := <-r.broadcast:
			r.broadcastMsg(message)
		}
	}
}
