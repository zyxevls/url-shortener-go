package usecase

import "time"

type URLUseCase interface {
	CreateShortURL(original string, alias string, expire *time.Time) (string, error)
	GetOriginalURL(code string) (string, error)
	IncrementClick(code string)
}
