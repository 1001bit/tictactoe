package server

import (
	"github.com/1001bit/tictactoe/services/game/hub"
	"github.com/go-chi/chi/v5"
)

func (s *Server) newRouter() *chi.Mux {
	r := chi.NewRouter()
	hub := hub.New()
	go hub.Run()
	go hub.DummyBroadcast()

	r.Get("/roomsSSE", hub.HandleSSE)

	return r
}
