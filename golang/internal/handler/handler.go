package handler

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	pb "github.com/alinux78/ulrshortener/internal/service/api/proto"
	"google.golang.org/grpc"
)

type UrlShortenerHadler interface {
	Shorten(w http.ResponseWriter, r *http.Request)
	Resolve(w http.ResponseWriter, r *http.Request)
	Redirect(short_url string, w http.ResponseWriter, r *http.Request)
}
type uRLShortener struct {
}

type fullUrl struct {
	URL string `json:"url"`
}

type shortenedUrl struct {
	ShortURL string `json:"short_url"`
}

var grpcClient pb.UrlShortenerServiceClient
var grpcConn *grpc.ClientConn

func initGrpcClient() {
	grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//TODO close connection
	grpcClient = pb.NewUrlShortenerServiceClient(grpcConn)
}

func NewURLShortener() UrlShortenerHadler {
	initGrpcClient()
	return &uRLShortener{}
}

func (h *uRLShortener) Stop() {
	grpcConn.Close()
}

func (h *uRLShortener) Shorten(w http.ResponseWriter, r *http.Request) {
	var req fullUrl
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	_, err := url.Parse(req.URL)
	if _, err := url.Parse(req.URL); err != nil {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	grpcResp, err := grpcClient.Shorten(ctx, &pb.UrlShortenRequest{Url: req.URL})
	if err != nil {
		slog.Error("grpc error", slog.String("error", err.Error()))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := shortenedUrl{ShortURL: getFullUrl(grpcResp.ShortUrl, r)}
	json.NewEncoder(w).Encode(resp)
}

func (h *uRLShortener) Redirect(shortURL string, w http.ResponseWriter, r *http.Request) {
	url, err := h.resolveShortURL(shortURL)
	if err != nil {
		slog.Error("grpc error", slog.String("error", err.Error()))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if url == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	url, _ = addSchemeIfMissing(url)
	slog.Debug("redirecting", "url", url)
	http.Redirect(w, r, url, http.StatusMovedPermanently)

}

func (h *uRLShortener) Resolve(w http.ResponseWriter, r *http.Request) {
	var req shortenedUrl
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	url, err := h.resolveShortURL(req.ShortURL)
	if err != nil {
		slog.Error("grpc error", slog.String("error", err.Error()))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if url == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	resp := fullUrl{URL: url}
	json.NewEncoder(w).Encode(resp)
}

func (h *uRLShortener) resolveShortURL(shortURL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	grpcResp, err := grpcClient.Resolve(ctx, &pb.UrlResolveRequest{ShortUrl: shortURL})
	if err != nil {
		slog.Error("grpc error", slog.String("error", err.Error()))
		return "", err
	}
	return grpcResp.Url, nil
}

func getFullUrl(url string, r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	host := r.Host
	fullURL := scheme + "://" + host + "/" + url

	return fullURL
}
func addSchemeIfMissing(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "http"
	}

	return parsedURL.String(), nil
}
