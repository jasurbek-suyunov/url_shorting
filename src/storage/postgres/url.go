package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/jmoiron/sqlx"
)

type urlRepo struct {
	db *sqlx.DB
}

const (
	urlTable  = "urls"
	urlFields = `id,user_id, org_path, short_path, counter, created_at, updated_at, status, qr_code_path`
)

// CreateUrl implements repository.UrlI
func (u *urlRepo) CreateUrl(ctx context.Context, url *models.Url) (*models.Url, error) {

	// response object
	var result models.Url

	// query
	query := `INSERT INTO urls(org_path, short_path, counter, created_at, updated_at, user_id, qr_code_path,status) VALUES($1, $2, $3, $4, $5, $6, $7,$8) RETURNING ` + urlFields

	// exec and scan
	err := u.db.QueryRow(
		query,
		url.OrgPath,
		url.ShortPath,
		url.Counter,
		url.CreatedAt,
		url.UpdatedAt,
		url.UserID,
		url.QrCodePath,
		url.Status,
	).Scan(
		&result.ID,
		&result.UserID,
		&result.OrgPath,
		&result.ShortPath,
		&result.Counter,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.Status,
		&result.QrCodePath,
	)

	// check error
	if err != nil {
		log.Printf("Method: CreateUrl, Error: %v", err)
		return nil, err
	}

	// return  if success
	return &result, nil
}

// DeleteUrlByShortURL implements storage.UrlI
func (u *urlRepo) DeleteUrlByShortURL(ctx context.Context, short_path string) error {

	// query
	query := `DELETE FROM urls WHERE short_path = $1`

	// exec
	result, err := u.db.Exec(query, short_path)

	// check error
	if err != nil {
		log.Printf("Method: DeleteUrlByShortURL, Error: %v", err)
		return err
	}

	// check affected rows
	if rows, err := result.RowsAffected(); err != nil || rows == 0 {
		log.Printf("Method: DeleteUrlByShortURL, Error: %v", err)
		return sql.ErrNoRows
	}

	// return nil if success
	return nil
}

// DeleteUrlByID implements storage.UrlI
func (u *urlRepo) DeleteUrlByID(ctx context.Context, id string) error {

	// query
	query := `DELETE FROM urls WHERE id = $1`

	// exec
	result, err := u.db.Exec(query, id)

	// check error
	if err != nil {
		log.Printf("Method: DeleteUrlByID, Error: %v", err)
		return err
	}

	// check affected rows
	if rows, err := result.RowsAffected(); err != nil || rows == 0 {
		log.Printf("Method: DeleteUrlByID, Error: %v", err)
		return sql.ErrNoRows
	}

	// return nil if success
	return nil
}

// GetUrlByShortPath implements storage.UrlI
func (u *urlRepo) GetUrlByShortPath(ctx context.Context, shortPath string) (*models.Url, error) {

	var result models.Url
	query := `UPDATE urls SET counter=counter+1 WHERE short_path = $1 RETURNING ` + urlFields

	err := u.db.QueryRow(
		query,
		shortPath,
	).Scan(
		&result.ID,
		&result.UserID,
		&result.OrgPath,
		&result.ShortPath,
		&result.Counter,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.Status,
		&result.QrCodePath,
	)

	if err != nil {
		log.Printf("Method: GetUrlByShortPath, Error: %v", err)
		return nil, err
	}

	return &result, nil

}

// GetUrlByID implements storage.UrlI
func (u *urlRepo) GetUrlByID(ctx context.Context, UserID string) (*models.Url, error) {

	// response object
	var url models.Url

	// query
	query := `SELECT ` + urlFields + ` FROM ` + urlTable + ` WHERE id = $1`

	// exec and scan
	err := u.db.QueryRow(
		query,
		UserID,
	).Scan(
		&url.ID,
		&url.UserID,
		&url.OrgPath,
		&url.ShortPath,
		&url.Counter,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.Status,
		&url.QrCodePath,
	)

	// check error
	if err != nil {
		log.Printf("Method: GetUrlByID, Error: %v", err)
		return nil, err
	}
	return &url, nil
}

// GetUrls implements storage.UrlI
func (u *urlRepo) GetUrls(ctx context.Context, id string) ([]*models.Url, error) {

	// response object
	urls := []*models.Url{}

	// query
	query := `SELECT ` + urlFields + ` FROM ` + urlTable + ` WHERE user_id = $1 `

	// exec
	rows, err := u.db.Query(query, id)

	// check error
	if err != nil {
		log.Printf("Method: GetUrls, Error: %v", err)
		return nil, err
	}

	// close rows
	defer rows.Close()

	// loop rows
	for rows.Next() {

		// append object
		var url models.Url

		// scan
		err := rows.Scan(
			&url.ID,
			&url.UserID,
			&url.OrgPath,
			&url.ShortPath,
			&url.Counter,
			&url.CreatedAt,
			&url.UpdatedAt,
			&url.Status,
			&url.QrCodePath,
		)

		// check error
		if err != nil {
			log.Printf("Method: GetUrls, Error: %v", err)
			return nil, err
		}

		// append to response object
		urls = append(urls, &url)
	}

	// return nil if success
	return urls, nil
}

// UpdateUrl implements storage.UrlI
func (u *urlRepo) UpdateUrl(ctx context.Context, url *models.Url) (*models.Url, error) {

	// response object
	var upt models.Url

	// query
	query := "UPDATE urls SET org_path = $1 ,updated_at= $2 WHERE id = $3"

	// exec and scan
	err := u.db.QueryRow(query, url.OrgPath, url.UpdatedAt, url.ID).Scan(
		&upt.ID,
		&upt.UserID,
		&upt.OrgPath,
		&upt.ShortPath,
		&upt.Counter,
		&upt.CreatedAt,
		&upt.UpdatedAt,
		&upt.Status,
	)

	// check error
	if err != nil {
		log.Printf("Method: UpdateUrl, Error: %v", err)
		return nil, err
	}

	// return if success
	return &upt, nil
}

func NewUrlRepo(db *sqlx.DB) *urlRepo {
	return &urlRepo{db}
}
