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
	urlTable  = "urls"
	urlFields = ` id, user_id, org_path, short_path, counter, created_at,updated_at,status `
)

// DeleteUrl implements storage.UrlI
func (u *urlRepo) DeleteUrl(ctx context.Context, urlID string) (*models.Url, error) {

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, urlTable)
	_, err := u.db.Exec(query, urlID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetUrlByID implements storage.UrlI
func (u *urlRepo) GetUrlByID(ctx context.Context, urlID string) (*models.Url, error) {
	url := &models.Url{}
	query := fmt.Sprintf(`
	SELECT
	id, 
	user_id,
	org_path, 
	short_path,
	counter, 
	created_at,
	update_at,
	type
	FROM
		%s
	WHERE
		id = $1`,
		userTable)

	row := u.db.QueryRow(query, urlID)
	err := row.Scan(
		&url.ID,
		&url.UserID,
		&url.OrgPath,
		&url.ShortPath,
		&url.Counter,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.Status,
	)
	if err != nil {
		return nil, err
	}
	return url, nil
}

// GetUrls implements storage.UrlI
func (u *urlRepo) GetUrls(ctx context.Context, url string) (*models.GetAllUrl, error) {
	// var urls *models.GetAllUrl
	// var count int = 0
	// query := fmt.Sprintf(`
	// SELECT
	// id,
	// user_id,
	// org_path,
	// short_path,
	// counter,
	// created_at,
	// type
	// FROM
	// 	%s
	// `, urlTable)

	// rows, err := u.db.Query(query)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// for rows.Next() {

	// 	err := rows.Scan(
	// 		&url.Urls.ID,
	// 		&url.Urls.UserID,
	// 		&url.Urls.OrgPath,
	// 		&url.Urls.ShortPath,
	// 		&url.Urls.Counter,
	// 		&url.Urls.CreatedAt,
	// 		&url.Urls.Type,
	// 	)

	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	urls = append(urls, url)
	// }
	// return urls, nil
	return nil, nil
}

// UpdateUrl implements storage.UrlI
func (u *urlRepo) UpdateUrl(ctx context.Context, url *models.Url) (*models.Url, error) {
	upt := models.Url{}
	query := fmt.Sprintf("UPDATE %s SET short_path = $1 updated_at=$2 WHERE id = $3", urlTable)
	rows := u.db.QueryRow(query, url.ShortPath, url.UpdatedAt, url.ID)
	err := rows.Scan(
		&upt.ID,
		&upt.UserID,
		&upt.OrgPath,
		&upt.ShortPath,
		&upt.Counter,
		&upt.CreatedAt,
		&upt.UpdatedAt,
		&upt.Status,
	)
	if err != nil {
		return nil, err
	}
	return &upt, nil
}

// CreateUrl implements repository.UrlI
func (u *urlRepo) CreateUrl(ctx context.Context, url *models.Url) (*models.Url, error) {
	resp := models.Url{}
	// ...1: Creating url

	query := `INSERT INTO urls( status,org_path, short_path, counter, created_at, updated_at, user_id) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING ` + urlFields
	fmt.Println(query)
	if err := u.db.QueryRow(
		query,
		url.Status,
		url.OrgPath,
		url.ShortPath,
		url.Counter,
		url.CreatedAt,
		url.UpdatedAt,
		url.UserID,
	).Scan(
		&resp.ID,
		&resp.UserID,
		&resp.OrgPath,
		&resp.ShortPath,
		&resp.Counter,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.Status,
	); err != nil {
		fmt.Println(err)
		return nil, err
	}

	// ...2: Returning successful response
	return &resp, nil
}

func NewUrlRepo(db *sqlx.DB) *urlRepo {
	return &urlRepo{db}
}
