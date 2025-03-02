package server

import (
	"github.com/1001bit/tictactoe/services/game/hub"
	"github.com/1001bit/tictactoe/services/game/roomhub"
	"github.com/go-chi/chi/v5"
)

func (s *Server) newRouter() *chi.Mux {
	r := chi.NewRouter()
	hub := hub.New()
	go hub.Run()
	go hub.DummyBroadcast()

	roomHub := roomhub.New()
	go roomHub.Run()

	r.Get("/roomsSSE", hub.HandleSSE)
	r.Get("/roomWS/{roomID}", roomHub.HandleWS)

	return r
}
