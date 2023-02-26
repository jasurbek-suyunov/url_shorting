package postgres

import (
	"fmt"

	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}
	var (
		userTable = "users"
		id string
	)
// CreateUser implements repository.UserI
func (u *userRepo) CreateUser(createusermodel models.User) (models.User, error) {
		// ...0: Checking if user already exists
		query0 := fmt.Sprintf(`SELECT id FROM %s WHERE email = $1`, userTable)
		row := u.db.QueryRow(query0, createusermodel.Username)
		err := row.Scan(&id)
		if err == nil {
			return models.User{}, fmt.Errorf("username already authorized")
		}
	
		query1 := fmt.Sprintf(`SELECT id FROM %s WHERE phone = $1`, userTable)
		row1 := ar.db.QueryRow(query1, createusermodel.Email)
		err = row1.Scan(&id)
		if err == nil {
			return models.User{}, fmt.Errorf("phone already authorized")
		}

		query := fmt.Sprintf(
			`INSERT INTO
					 %s
				 (
					 phone,
					 email,
					 avatar_url,
					password_hash,
					type,
					status,
					language_id,
					created_at,
					updated_at,
					deleted_at
				 ) VALUES (
					 $1,
					 $2,
					 $3,
					 $4,
					$5,
					$6,
					$7,
					$8,
					$9,
					$10
				 ) RETURNING id, phone, email, type, language_id, status, created_at
			 `, usersTable)
	
		if err := ar.db.QueryRow(
			query,
			createusermodel.Phone,
			createusermodel.Email,
			createusermodel.AvatarUrl,
			createusermodel.Password,
			createusermodel.Type,
			createusermodel.Status,
			createusermodel.LanguageID,
			createusermodel.CreatedAt,
			createusermodel.UpdatedAt,
			createusermodel.DeletedAt,
		).Scan(
			&resp.ID,
			&resp.Phone,
			&resp.Email,
			&resp.Type,
			&resp.LanguageID,
			&resp.Status,
			&resp.CreatedAt,
		); err != nil {
			return resp, err
		}
	
		// ...2: Returning successful response
		return resp, nil
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db}
}
