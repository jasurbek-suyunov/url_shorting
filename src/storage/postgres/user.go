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

// DeleteUser implements storage.UserI
func (*userRepo) DeleteUser(ctx context.Context, user *models.User) (*models.User, error) {
	panic("unimplemented")
}

// GetUserByID implements storage.UserI
func (*userRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}

// GetUserByUsername implements storage.UserI
func (*userRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements storage.UserI
func (*userRepo) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	panic("unimplemented")
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

	resp := models.User{}
	// ...1: Creating user
	query := fmt.Sprintf(`INSERT INTO users(username, first_name, last_name, email, password_hash, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING %s`, userFields)

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
		fmt.Println("err: ", err.Error())
		return nil, err
	}

	// ...2: Returning successful response
	return &resp, nil
}
