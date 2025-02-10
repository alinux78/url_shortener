package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"log/slog"
	"net"

	pb "github.com/alinux78/ulrshortener/internal/service/api/proto"
	"google.golang.org/grpc"

	"github.com/alinux78/ulrshortener/internal/repository"
)

type URLShortener interface {
	Shorten(url string) (string, error)
	Resolve(url string) (string, error)
}

type uRLShortener struct {
	repo repository.Repository
}

func NewURLShortener(repo repository.Repository) URLShortener {
	return &uRLShortener{repo: repo}
}

func (s *uRLShortener) Shorten(url string) (string, error) {
	sum := sha256.Sum256([]byte(url))
	return base64.URLEncoding.EncodeToString(sum[:])[:6], nil
}

func (s *uRLShortener) Resolve(url string) (string, error) {
	return "", nil
}

type server struct {
	shortener uRLShortener
	pb.UnimplementedUrlShortenerServiceServer
}

func (s *server) Shorten(ctx context.Context, in *pb.UrlShortenRequest) (*pb.UrlShortenResponse, error) {
	enc, err := s.shortener.Shorten(in.Url)
	if err != nil {
		return nil, err
	}
	return &pb.UrlShortenResponse{ShortUrl: enc}, nil
}

func Start(repo repository.Repository) {
	port := 50051
	slog.Info("GRPC service server started on ", slog.Int("port", port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	uRLShortener := uRLShortener{repo: repo}
	s := grpc.NewServer()
	pb.RegisterUrlShortenerServiceServer(s, &server{shortener: uRLShortener})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
