package roomhub

type ClientRegRequest struct {
	roomID string
	client *Client
}

type RoomHub struct {
	rooms          map[string]*Room
	roomUnregister chan string

	clientRegister chan ClientRegRequest
}

func New() *RoomHub {
	return &RoomHub{
		rooms:          make(map[string]*Room),
		roomUnregister: make(chan string),

		clientRegister: make(chan ClientRegRequest),
	}
}

func (rh *RoomHub) Run() {
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
		}
	}
}
