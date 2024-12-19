package main

import (
	"log/slog"

	"github.com/alinux78/ulrshortener/internal/api"
	"github.com/alinux78/ulrshortener/internal/repository"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	port := 8080
	repo := repository.NewInMemoryRepository()
	api.Serve(port, repo)
}