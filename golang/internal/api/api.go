package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/alinux78/ulrshortener/internal/handler"
	"github.com/alinux78/ulrshortener/internal/repository"
	"github.com/alinux78/ulrshortener/internal/service"
	"github.com/gin-gonic/gin"
)

func Serve(port int, repo repository.Repository) {
	slog.Info("server started", slog.Int("port", port))
	service := service.NewURLShortener(repo)
	handler := handler.NewURLShortener(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.HandleMethodNotAllowed = true

	r.Use(gin.Logger())
	r.Use(auth()) //TODO use gin.BasicAuth

	r.POST("/shorten", func(c *gin.Context) {
		handler.Shorten(c.Writer, c.Request)
	})
	r.POST("/resolve", func(c *gin.Context) {
		handler.Resolve(c.Writer, c.Request)
	})

	middleware := r

	err := http.ListenAndServe((":" + strconv.Itoa(port)), middleware)
	if err != nil {
		slog.Error("error starting server", slog.String("error", err.Error()))
	}
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Debug("skip auth for request", "method", c.Request.Method, "path", c.Request.RequestURI)
		c.Next()
	}
}
