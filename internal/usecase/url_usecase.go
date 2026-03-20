package usecase

import (
	"errors"
	"time"

	"github.com/zyxevls/internal/domain"
	"github.com/zyxevls/pkg/utils"
)

type URLRepository interface {
	Save(url *domain.URL) error
	FindByCode(code string) (*domain.URL, error)
	IncrementClick(code string) error
}

type CacheRepository interface {
	Set(code string, original string, ttl time.Duration) error
	Get(code string) (string, error)
	IncrementClick(code string)
}

type URLUseCase struct {
	repo  URLRepository
	cache CacheRepository
}

func NewUrlUseCase(r URLRepository, c CacheRepository) *URLUseCase {
	return &URLUseCase{repo: r, cache: c}
}

func (u *URLUseCase) CreateShortURL(original string, alias string, expire *time.Time) (*domain.URL, error) {
	code := alias
	if code == "" {
		code = utils.GenerateShortCode(6)
	}

	var expiresAt *time.Time
	if expire != nil {
		expiresAt = expire
	}

	url := &domain.URL{
		OriginalURL: original,
		ShortCode:   code,
		CustomAlias: alias,
		ExpiresAt:   expiresAt,
		CreatedAt:   time.Now(),
	}

	err := u.repo.Save(url)
	if err != nil {
		return nil, err
	}

	//cache
	ttl := time.Hour * 24
	if expiresAt != nil {
		ttl = time.Until(*expire)
	}

	_ = u.cache.Set(code, original, ttl)

	return url, nil
}

func (u *URLUseCase) GetOriginalURL(code string) (string, error) {
	//cek redis
	cached, err := u.cache.Get(code)
	if err == nil {
		go u.cache.IncrementClick(code)
		return cached, nil
	}

	//cek db
	url, err := u.repo.FindByCode(code)
	if err != nil {
		return "", err
	}

	if url.ExpiresAt != nil && time.Now().After(*url.ExpiresAt) {
		return "", errors.New("Link expired")
	}

	//simpan ke redis
	ttl := time.Hour * 24
	if url.ExpiresAt != nil {
		ttl = time.Until(*url.ExpiresAt)
	}

	_ = u.cache.Set(code, url.OriginalURL, ttl)

	go u.repo.IncrementClick(code)

	return url.OriginalURL, nil

}

func (u *URLUseCase) IncrementClick(code string) {
	_ = u.repo.IncrementClick(code)
	u.cache.IncrementClick(code)
}
