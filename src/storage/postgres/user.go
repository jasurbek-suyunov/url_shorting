package postgres

import (
	"context"
	"fmt"

	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

const (
	userTable  = "users"
	userFields = `id, username, first_name, last_name, email, password_hash, created_at`
)

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db}
}

// CreateUser implements repository.UserI
func (u *userRepo) CreateUser(ctx context.Context, usr *models.User) (*models.User, error) {

	resp := &models.User{}
	// ...1: Creating user
	query := fmt.Sprintf(
		`INSERT INTO
					 %s
				 (
					 username,
					 first_name,
					 last_name,
					 email,
					 password_hash,
					 created_at
				 ) VALUES (
					 $1,
					 $2,
					 $3,
					 $4,
					 $5,
					 $6
				 ) RETURNING %s
			 `, userTable, userFields)

	if err := u.db.QueryRow(
		query,
		usr.Username,
		usr.FirstName,
		usr.LastName,
		usr.Email,
		usr.PasswordHash,
		usr.CreatedAt,
	).Scan(
		&resp.ID,
		&resp.Username,
		&resp.FirstName,
		&resp.LastName,
		&resp.Email,
		&resp.PasswordHash,
		&resp.CreatedAt,
	); err != nil {
		return nil, err
	}

	// ...2: Returning successful response
	return resp, nil
}
