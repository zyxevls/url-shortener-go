package domain

import "time"

type URL struct {
	ID          string
	OriginalURL string
	ShortCode   string
	CustomAlias string
	ClickCount  int
	ExpiresAt   *time.Time
	CratedAt    time.Time
}
