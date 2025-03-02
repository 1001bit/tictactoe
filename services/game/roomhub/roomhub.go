package roomhub

import "github.com/1001bit/tictactoe/services/game/hub"

type ClientRegRequest struct {
	roomID string
	client *Client
}

type RoomHub struct {
	rooms          map[string]*Room
	roomUnregister chan string

	clientRegister chan ClientRegRequest

	roomsUpdateChan chan []struct{}
}

func New() *RoomHub {
	return &RoomHub{
		rooms:          make(map[string]*Room),
		roomUnregister: make(chan string),

		clientRegister: make(chan ClientRegRequest),

		roomsUpdateChan: make(chan []struct{}),
	}
}

func (rh *RoomHub) GenerateAndBroadcastRoomsMsg(h *hub.Hub) {
	roomsMsg := make([]hub.RoomMsg, 0, len(rh.rooms))
	for _, room := range rh.rooms {
		roomsMsg = append(roomsMsg, hub.RoomMsg{
			Id:      room.id,
			Players: len(room.clients),
		})
	}

	h.BroadcastRoomsMsg(roomsMsg)
}

func (rh *RoomHub) Run(h *hub.Hub) {
	for {
		select {
		case req := <-rh.clientRegister:
			room, ok := rh.rooms[req.roomID]
			if !ok {
				room = NewRoom(req.roomID)
				go room.Run(rh)
				rh.rooms[req.roomID] = room
			}
			req.client.room = room
			room.register <- req.client
		case roomId := <-rh.roomUnregister:
			room, ok := rh.rooms[roomId]
			if !ok {
				continue
			}
			close(room.broadcast)
			close(room.register)
			close(room.unregister)
			delete(rh.rooms, roomId)

			rh.GenerateAndBroadcastRoomsMsg(h)
		case <-rh.roomsUpdateChan:
			rh.GenerateAndBroadcastRoomsMsg(h)
		}
	}
}
