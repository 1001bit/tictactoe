package roomhub

import (
	"encoding/json"
	"log/slog"
)

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

func (rh *RoomHub) roomsNotify(notifyChan chan<- []byte) {
	type RoomNotify struct {
		Rooms []string `json:"rooms"`
	}

	rn := RoomNotify{
		Rooms: make([]string, 0, len(rh.rooms)),
	}

	for roomId := range rh.rooms {
		rn.Rooms = append(rn.Rooms, roomId)
	}

	b, err := json.Marshal(rn)
	if err != nil {
		slog.Error("error marshaling", "err", err.Error())
		return
	}

	notifyChan <- b
}

func (rh *RoomHub) Run(notifyChan chan<- []byte) {
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

			rh.roomsNotify(notifyChan)
		case roomId := <-rh.roomUnregister:
			room, ok := rh.rooms[roomId]
			if !ok {
				continue
			}
			close(room.broadcast)
			close(room.register)
			close(room.unregister)
			delete(rh.rooms, roomId)

			rh.roomsNotify(notifyChan)
		}
	}
}
