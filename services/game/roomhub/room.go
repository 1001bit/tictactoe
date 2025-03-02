package roomhub

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

func (r *Room) Run(roomHub *RoomHub) {
	defer func() {
		roomHub.roomUnregister <- r.id
	}()

	slog.Info("Room started", "id", r.id)
	for {
		select {
		case client, ok := <-r.register:
			if !ok {
				return
			}
			if len(r.clients) >= 2 {
				close(client.sendCh)
				continue
			}
			r.clients[client] = true
			roomHub.roomsUpdateChan <- nil
		case client := <-r.unregister:
			if _, ok := r.clients[client]; !ok {
				continue
			}
			delete(r.clients, client)
			close(client.sendCh)

			if len(r.clients) == 0 {
				return
			} else {
				roomHub.roomsUpdateChan <- nil
			}
		case message := <-r.broadcast:
			for client := range r.clients {
				select {
				case client.sendCh <- message:
				default:
					delete(r.clients, client)
					close(client.sendCh)
				}
			}
		}
	}
}
