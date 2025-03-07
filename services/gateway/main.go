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

	port := os.Getenv("PORT")
	gameAddr := "http://game:" + os.Getenv("PORT")
	if err := s.Run(port, gameAddr); err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Server stopped")
}
