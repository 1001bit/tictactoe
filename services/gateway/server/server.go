package server

import (
	"log/slog"
	"net/http"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run(port, gameAddr string) error {
	addr := ":" + port
	slog.Info(
		"Start server",
		"addr", addr,
	)

	r, err := s.newRouter(gameAddr)
	if err != nil {
		return err
	}
	http.ListenAndServe(addr, r)

	return nil
}
