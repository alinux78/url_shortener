package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/alinux78/ulrshortener/internal/handler"
	"github.com/gin-gonic/gin"
)

func Serve(port int, handler handler.UrlShortenerHadler) {

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

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/:short_url", func(c *gin.Context) {
		shortUrl := c.Param("short_url")
		slog.Debug("redirecting", "short_url", shortUrl)
		handler.Redirect(shortUrl, c.Writer, c.Request)
	})

	middleware := r

	err := http.ListenAndServe((":" + strconv.Itoa(port)), middleware)
	if err != nil {
		slog.Error("error starting server", slog.String("error", err.Error()))
	}
	slog.Info("server started", slog.Int("port", port))

}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Debug("skip auth for request", "method", c.Request.Method, "path", c.Request.RequestURI)
		c.Next()
	}
}
