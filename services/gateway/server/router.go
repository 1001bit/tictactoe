package server

import (
	"net/http"

	"github.com/1001bit/tictactoe/services/gateway/httpproxy"
	"github.com/1001bit/tictactoe/services/gateway/server/handler"
	"github.com/go-chi/chi/v5"
)

func (s *Server) newRouter(gameAddr string) (*chi.Mux, error) {
	r := chi.NewRouter()

	gameServiceProxy, err := httpproxy.New(gameAddr)
	if err != nil {
		return nil, err
	}

	r.Get("/", handler.HandleHome)
	r.Get("/static/*", http.StripPrefix("/static", handler.Static()).ServeHTTP)
	r.Get("/api/game/*", gameServiceProxy.ProxyHandler("/api/game").ServeHTTP)

	return r, nil
}
