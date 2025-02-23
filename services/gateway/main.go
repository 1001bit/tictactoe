package main

import (
	"log/slog"
	"os"

	"github.com/1001bit/tictactoe/services/gateway/server"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func main() {
	s := server.New()
	if err := s.Run(os.Getenv("PORT")); err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Server stopped")
}
