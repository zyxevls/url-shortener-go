package postgres

import (
	"fmt"
	"sync"

	"github.com/zyxevls/internal/domain"
)

type URLRepository struct {
	mu     sync.RWMutex
	byCode map[string]*domain.URL
}

func NewURLRepository() *URLRepository {
	return &URLRepository{
		byCode: make(map[string]*domain.URL),
	}
}

func (r *URLRepository) Save(url *domain.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.byCode[url.ShortCode]; exists {
		return fmt.Errorf("short code already exists")
	}

	r.byCode[url.ShortCode] = url
	return nil
}

func (r *URLRepository) FindByCode(code string) (*domain.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	url, ok := r.byCode[code]
	if !ok {
		return nil, fmt.Errorf("short code not found")
	}

	return url, nil
}

func (r *URLRepository) IncrementClick(code string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	url, ok := r.byCode[code]
	if !ok {
		return fmt.Errorf("short code not found")
	}

	url.ClickCount++
	return nil
}
