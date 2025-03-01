package handler

import (
	"log/slog"
	"net/http"
)

func HandleRoomsSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust for security

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("data: hi\n\n"))
	flusher.Flush()

	<-r.Context().Done()
	slog.Info("Connection closed")
}
