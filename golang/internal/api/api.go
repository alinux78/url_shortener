package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/alinux78/ulrshortener/internal/handler"
	"github.com/alinux78/ulrshortener/internal/repository"
	"github.com/alinux78/ulrshortener/internal/service"
	"github.com/go-chi/chi/v5"
)

func Serve(port int, repo repository.Repository) {
	slog.Info("server started", slog.Int("port", port))
	service := service.NewURLShortener(repo)
	handler := handler.NewURLShortener(service)

	// r := mux.NewRouter()
	// r.HandleFunc("/shorten", handler.Shorten).Methods("POST")
	// r.HandleFunc("/resolve", handler.Resolve).Methods("GET")
	// http.Handle("/", r)

	r := chi.NewRouter()

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	r.Post("/shorten", handler.Shorten)
	r.Post("/resolve", handler.Resolve)
	http.Handle("/", r)

	err := http.ListenAndServe((":" + strconv.Itoa(port)), nil)
	if err != nil {
		slog.Error("error starting server", slog.String("error", err.Error()))
	}
}
