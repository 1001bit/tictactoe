package hub

import (
	"fmt"
	"net/http"
)

type Client struct {
	sendCh chan []byte
}

func (hub *Hub) HandleSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	c := &Client{
		sendCh: make(chan []byte, 5),
	}

	hub.register <- c
	defer func() {
		hub.unregister <- c
	}()

	for {
		select {
		case message, ok := <-c.sendCh:
			if !ok {
				return
			}
			fmt.Fprintf(w, "data: %s\n\n", message)
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
