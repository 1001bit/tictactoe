package hub

import (
	"encoding/json"
	"log/slog"
)

type RoomMsg struct {
	Id      string `json:"id"`
	Players int    `json:"players"`
}

type Hub struct {
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool

	roomsMsg []byte
}

func New() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),

		roomsMsg: []byte(`{"rooms": []}`),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			select {
			case client.sendCh <- h.roomsMsg:
			default:
				delete(h.clients, client)
				close(client.sendCh)
			}
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

func (h *Hub) BroadcastRoomsMsg(roomsMsg []RoomMsg) {
	msgByte, err := json.Marshal(map[string][]RoomMsg{"rooms": roomsMsg})
	if err != nil {
		slog.Error("error marshaling", "err", err.Error())
		return
	}

	h.roomsMsg = msgByte
	h.broadcast <- h.roomsMsg
}
