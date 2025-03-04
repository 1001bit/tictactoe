package room

import "github.com/1001bit/tictactoe/services/game/hub"

type ClientRegRequest struct {
	roomID string
	client *Client
}

type RoomStore struct {
	rooms          map[string]*Room
	roomUnregister chan string

	clientRegister chan ClientRegRequest

	roomsUpdateChan chan []struct{}
}

func NewStore() *RoomStore {
	return &RoomStore{
		rooms:          make(map[string]*Room),
		roomUnregister: make(chan string),

		clientRegister: make(chan ClientRegRequest),

		roomsUpdateChan: make(chan []struct{}),
	}
}

func (rs *RoomStore) GenerateAndBroadcastRoomsMsg(h *hub.Hub) {
	roomsMsg := make([]hub.RoomMsg, 0, len(rs.rooms))
	for _, room := range rs.rooms {
		roomsMsg = append(roomsMsg, hub.RoomMsg{
			Id:      room.id,
			Players: len(room.clients),
		})
	}

	h.BroadcastRoomsMsg(roomsMsg)
}

func (rs *RoomStore) Run(h *hub.Hub) {
	for {
		select {
		case req := <-rs.clientRegister:
			room, ok := rs.rooms[req.roomID]
			if !ok {
				room = NewRoom(req.roomID)
				go room.Run(rs)
				rs.rooms[req.roomID] = room
			}
			req.client.room = room
			room.register <- req.client
		case roomId := <-rs.roomUnregister:
			room, ok := rs.rooms[roomId]
			if !ok {
				continue
			}
			close(room.broadcast)
			close(room.register)
			close(room.unregister)
			delete(rs.rooms, roomId)

			rs.GenerateAndBroadcastRoomsMsg(h)
		case <-rs.roomsUpdateChan:
			rs.GenerateAndBroadcastRoomsMsg(h)
		}
	}
}
