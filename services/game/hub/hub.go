package hub

import "log/slog"

type Hub struct {
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
}

func New() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			slog.Info("Hub Client registered")
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.sendCh)
				slog.Info("Hub Client unregistered")
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.sendCh <- message:
				default:
					delete(h.clients, client)
					close(client.sendCh)
				}
			}
		}
	}
}

func (h *Hub) BroadcastIn() chan<- []byte {
	return h.broadcast
}
