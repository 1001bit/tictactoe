package server

import (
	"github.com/1001bit/tictactoe/services/game/server/handler"
	"github.com/go-chi/chi/v5"
)

func (s *Server) newRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/roomsSSE", handler.HandleRoomsSSE)

	return r
}
