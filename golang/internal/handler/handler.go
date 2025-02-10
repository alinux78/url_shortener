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

type uRLShortener struct {
}

type urlShortenRequest struct {
	URL string `json:"url"`
}

type urlShortenResponse struct {
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

func NewURLShortener() *uRLShortener {
	initGrpcClient()
	return &uRLShortener{}
}

func (h *uRLShortener) Stop() {
	grpcConn.Close()
}

func (h *uRLShortener) Shorten(w http.ResponseWriter, r *http.Request) {
	var req urlShortenRequest
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

	resp := urlShortenResponse{ShortURL: grpcResp.ShortUrl}
	json.NewEncoder(w).Encode(resp)
}

func (h *uRLShortener) Resolve(w http.ResponseWriter, r *http.Request) {
	//send not implemented
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
