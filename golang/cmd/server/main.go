package main

import (
	"log/slog"

	"github.com/alinux78/ulrshortener/internal/api"
	"github.com/alinux78/ulrshortener/internal/repository"
	"github.com/alinux78/ulrshortener/internal/service"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	port := 8080
	repo := repository.NewInMemoryRepository()
	go service.Start(repo)
	api.Serve(port)
}
