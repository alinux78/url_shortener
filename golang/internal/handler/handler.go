package handler

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"

	pb "github.com/alinux78/ulrshortener/internal/service/api/proto"
	"google.golang.org/grpc"
)

type UrlShortenerHadler interface {
	Shorten(w http.ResponseWriter, r *http.Request)
	Resolve(w http.ResponseWriter, r *http.Request)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	grpcResp, err := grpcClient.Shorten(ctx, &pb.UrlShortenRequest{Url: req.URL})
	if err != nil {
		slog.Error("grpc error", slog.String("error", err.Error()))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := shortenedUrl{ShortURL: grpcResp.ShortUrl}
	json.NewEncoder(w).Encode(resp)
}

func (h *uRLShortener) Resolve(w http.ResponseWriter, r *http.Request) {
	var req shortenedUrl
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	grpcResp, err := grpcClient.Resolve(ctx, &pb.UrlResolveRequest{ShortUrl: req.ShortURL})
	if err != nil {
		slog.Error("grpc error", slog.String("error", err.Error()))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if grpcResp.Url == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	resp := fullUrl{URL: grpcResp.Url}
	json.NewEncoder(w).Encode(resp)

}
