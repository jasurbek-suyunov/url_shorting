package postgres

import (
	"context"
	"database/sql"
	"log"

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

// CreateUser implements repository.UserI
func (u *userRepo) CreateUser(ctx context.Context, usr *models.User) (*models.User, error) {

	// response model
	resp := models.User{}

	// query
	query := `INSERT INTO users(username, first_name, last_name, email, password_hash, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING ` + userFields

	// exec and scan
	err := u.db.QueryRow(
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
	)

	// check if user exists
	if err != nil {
		log.Printf("Method: CreateUser, Error: %v", err)
		return nil, err
	}

	// return result if success
	return &resp, nil
}

// DeleteUser implements storage.UserI
func (u *userRepo) DeleteUser(ctx context.Context, uresID string) error {

	//query
	query := `DELETE FROM users WHERE id = $1`

	// exec
	result, err := u.db.ExecContext(ctx, query, uresID)

	// check if error
	if err != nil {
		log.Printf("Method: DeleteUser, Error: %v", err)
		return err
	}

	//check if user exists and affected count
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 || err != nil {
		return sql.ErrNoRows
	}

	// if success
	return nil
}

// GetUserByID implements storage.UserI
func (u *userRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {

	// response model
	var result models.User

	// query
	query := `SELECT ` + userFields + ` FROM users WHERE id = $1`

	// exec and scan
	if err := u.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&result.ID,
		&result.Username,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.PasswordHash,
		&result.CreatedAt,
	); err != nil {
		log.Printf("Method: GetUserByID, Error: %v", err)
		return nil, err
	}

	// return result
	return &result, nil
}

// GetUserByUsername implements storage.UserI
func (u *userRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {

	// response model
	var result models.User

	// query
	query := `SELECT ` + userFields + ` FROM users WHERE username = $1`

	// exec and scan
	err := u.db.QueryRowContext(
		ctx,
		query,
		username,
	).Scan(
		&result.ID,
		&result.Username,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.PasswordHash,
		&result.CreatedAt,
	)

	// check error
	if err != nil {
		log.Printf("Method: GetUserByUsername, Error: %v", err)
		return nil, err
	}

	// return result
	return &result, nil
}

// UpdateUser implements storage.UserI
func (u *userRepo) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {

	// response model
	var result models.User

	// query
	query := `UPDATE users SET username = $1, first_name = $2, last_name = $3, email = $4 WHERE id = $5 RETURNING ` + userFields

	// exec and scan
	err := u.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.ID,
	).Scan(
		&result.ID,
		&result.Username,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.PasswordHash,
		&result.CreatedAt,
	)

	// check error
	if err != nil {
		log.Printf("Method: UpdateUser, Error: %v", err)
		return nil, err
	}

	// return result
	return &result, nil
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db}
}
