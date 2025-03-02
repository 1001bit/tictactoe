package roomhub

import (
	"bytes"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	maxMsgSize = 512
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = pongWait * 9 / 10
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	sendCh chan []byte
	room   *Room
}

func (c *Client) readPump(conn *websocket.Conn) {
	defer func() {
		conn.Close()
		if c.room != nil {
			c.room.unregister <- c
		} else {
			close(c.sendCh)
		}
	}()

	conn.SetReadLimit(maxMsgSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error("error reading", "err", err.Error())
			}
			break
		}

		msg = bytes.TrimSpace(bytes.Replace(msg, []byte{'\n'}, []byte{' '}, -1))
		// TODO: handle message instead of broadcast
		if c.room == nil {
			return
		}
		c.room.broadcast <- msg
	}
}

func (c *Client) writePump(conn *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.sendCh:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		case <-ticker.C:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (rh *RoomHub) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("error upgrading", "err", err.Error())
		return
	}

	c := &Client{
		sendCh: make(chan []byte, 9),
	}
	roomID := r.PathValue("roomID")

	req := ClientRegRequest{
		roomID: roomID,
		client: c,
	}
	rh.clientRegister <- req

	slog.Info("Client started")
	go c.readPump(conn)
	go c.writePump(conn)
}
