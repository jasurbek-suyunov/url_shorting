package postgres

import (
	"context"
	"fmt"

	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/jmoiron/sqlx"
)

type urlRepo struct {
	db *sqlx.DB
}

const (
	urlTable = "urls"
	urlFields = `id, user_id, org_path, short_path, counter, created_at, type`
)
// CreateUrl implements repository.UrlI
func (u *urlRepo) CreateUrl(ctx context.Context, url *models.Url) (*models.Url, error) {
	resp := &models.Url{}
	// ...1: Creating url
	query := fmt.Sprintf(
		`INSERT INTO
					 %s
				 (
					 user_id,
					 org_path,
					 short_path,
					 counter,
					created_at,
					type
				 ) VALUES (
					 $1,
					 $2,
					 $3,
					 $4,
					 $5
				 ) RETURNING %s
			 `,urlTable,urlFields)

	if err := u.db.QueryRow(
		query,
		url.ID,
		url.UserID,
		url.OrgPath,
		url.ShortPath,
		url.Counter,
		url.CreatedAt,
		url.Type,
	).Scan(
		&resp.ID,
		&resp.UserID,
		&resp.OrgPath,
		&resp.ShortPath,
		&resp.Counter,
		&resp.CreatedAt,
		&resp.Type,
	); err != nil {
		return nil, err
	}

	// ...2: Returning successful response
	return resp, nil
}

func NewUrlRepo(db *sqlx.DB) *urlRepo {
	return &urlRepo{db}
}
