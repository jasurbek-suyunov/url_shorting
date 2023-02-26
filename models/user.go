package models

type User struct {
	ID           string `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	CreatedAt    int64  `json:"created_at" db:"created_at"`
}

type UserSignUpRequest struct {
	Username        string `json:"username" binding:"required" db:"username"`
	FirstName       string `json:"first_name" binding:"required" db:"first_name"`
	LastName        string `json:"last_name" binding:"required" db:"last_name"`
	Email           string `json:"email" binding:"required" db:"email"`
	Password        string `json:"password" binding:"required" db:"password"`
	ConfirmPassword string `json:"confirm_password" binding:"required" db:"confirm_password"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}

type LoginResponse struct {
	Data  *UserResponse `json:"data"`
	Error string        `json:"error"`
	Code  int           `json:"code"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password"`
}
