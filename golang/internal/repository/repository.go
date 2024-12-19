package repository

type Repository interface {
	Save(shortURL string, longURL string) error
	Load(shortURL string) (string, bool, error)
}

type inMemoryRepository struct {
	urls map[string]string
}

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		urls: make(map[string]string),
	}
}

func (r *inMemoryRepository) Save(shortURL string, longURL string) error {
	r.urls[shortURL] = longURL
	return nil
}

func (r *inMemoryRepository) Load(shortURL string) (string, bool, error) {
	longURL, ok := r.urls[shortURL]
	return longURL, ok, nil
}
