package api

import (
	"log"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"testing"

	"github.com/alinux78/ulrshortener/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func findAvailablePort() (int, error) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func TestServe(t *testing.T) {

	mockHandler := new(mocks.UrlShortenerHadler)

	port, err := findAvailablePort()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	slog.Info("Listening on", slog.Int("port", port))

	go Serve(port, mockHandler)

	mockHandler.On("Shorten", mock.Anything, mock.Anything).Return()

	url := "http://localhost:" + strconv.Itoa(port)
	// Test /shorten endpoint
	resp, err := http.Post(url+"/shorten", "application/json", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockHandler.On("Resolve", mock.Anything, mock.Anything).Return()

	// Test /resolve endpoint
	resp, err = http.Post(url+"/resolve", "application/json", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	resp, err = http.Get(url + "/resolve")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)

	mockHandler.AssertExpectations(t)
}
