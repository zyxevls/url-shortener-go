package usecase

import (
	"time"

	"github.com/zyxevls/internal/domain"
	"github.com/zyxevls/pkg/utils"
)

type URLRepository interface {
	Save(url *domain.URL) error
	FindByCode(code string) (*domain.URL, error)
	IncrementClick(code string) error
}

type URLUseCase struct {
	repo URLRepository
}

func NewUrlUseCase(r URLRepository) *URLUseCase {
	return &URLUseCase{repo: r}
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

	return url, nil
}

func (u *URLUseCase) GetOriginalURL(code string) (string, error) {
	url, err := u.repo.FindByCode(code)
	if err != nil {
		return "", err
	}

	return url.OriginalURL, nil
}

func (u *URLUseCase) IncrementClick(code string) {
	_ = u.repo.IncrementClick(code)
}
