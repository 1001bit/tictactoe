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

func (s *Server) Run(port string) error {
	addr := ":" + port
	slog.Info(
		"Start server",
		"addr", addr,
	)

	http.ListenAndServe(addr, s.newRouter())

	return nil
}
