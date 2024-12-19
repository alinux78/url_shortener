package service

import (
	"crypto/sha256"
	"encoding/base64"

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
