package server

import (
	"net/http"

	"github.com/1001bit/tictactoe/services/gateway/server/handler"
	"github.com/go-chi/chi/v5"
)

func (s *Server) newRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handler.HandleHome)
	r.Get("/static/*", http.StripPrefix("/static", handler.Static()).ServeHTTP)

	return r
}
