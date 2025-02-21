package service

import (
	"testing"

	"github.com/alinux78/ulrshortener/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShorten(t *testing.T) {
	mockRepo := new(mocks.Repository)
	service := NewURLShortener(mockRepo)

	originalURL1 := "https://example1.com"
	mockRepo.On("Save", mock.Anything, originalURL1).Return(nil)

	shortURL1, err := service.Shorten(originalURL1)
	assert.NoError(t, err)
	assert.Equal(t, 6, len(shortURL1))

	originalURL2 := "https://example2.com"
	mockRepo.On("Save", mock.Anything, originalURL2).Return(nil)

	shortURL2, err := service.Shorten(originalURL2)
	assert.NoError(t, err)
	assert.Equal(t, 6, len(shortURL2))

	assert.NotEqual(t, shortURL1, shortURL2)

	mockRepo.AssertExpectations(t)
}

func TestResolve(t *testing.T) {
	mockRepo := new(mocks.Repository)
	service := NewURLShortener(mockRepo)

	shortURL := "abc123"
	expectedOriginalURL := "https://example.com"

	mockRepo.On("Load", shortURL).Return(expectedOriginalURL, true, nil)

	originalURL, err := service.Resolve(shortURL)
	assert.NoError(t, err)
	assert.Equal(t, expectedOriginalURL, originalURL)

	mockRepo.AssertExpectations(t)
}
