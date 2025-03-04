package server

import (
	"github.com/1001bit/tictactoe/services/game/hub"
	"github.com/1001bit/tictactoe/services/game/room"
	"github.com/go-chi/chi/v5"
)

func (s *Server) newRouter() *chi.Mux {
	r := chi.NewRouter()
	hub := hub.New()
	go hub.Run()

	roomStore := room.NewStore()
	go roomStore.Run(hub)

	r.Get("/roomsSSE", hub.HandleSSE)
	r.Get("/roomWS/{roomID}", roomStore.HandleWS)

	return r
}
