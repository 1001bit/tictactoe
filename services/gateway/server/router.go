package server

import (
	"github.com/1001bit/tictactoe/services/gateway/server/handler"
	"github.com/go-chi/chi/v5"
)

func (s *Server) newRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handler.Home)

	return r
}
