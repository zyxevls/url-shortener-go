package postgres

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/zyxevls/internal/domain"
)

type URLRepository struct {
	db *sqlx.DB
}

func NewURLRepository(db *sqlx.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Save(url *domain.URL) error {
	query := `
	INSERT INTO urls (orignal_url, short_code, custom_alias, expires_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`

	err := r.db.QueryRow(
		query,
		url.OriginalURL,
		url.ShortCode,
		url.CustomAlias,
		url.ExpiresAt,
	).Scan(&url.ID)

	return err
}

func (r *URLRepository) FindByCode(code string) (*domain.URL, error) {
	var url domain.URL

	query := `
	SELECT id, original_url, short_code, custom_alias, click_count, expires_at, created_at
	FROM urls
	WHERE short_code = $1
	`

	err := r.db.Get(&url, query, code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return &url, nil
}

func (r *URLRepository) IncrementClick(code string) error {
	query := `
	UPDATE urls
	SET click_count = click_count + 1
	WHERE short_code = $1
	`

	_, err := r.db.Exec(query, code)
	return err
}
