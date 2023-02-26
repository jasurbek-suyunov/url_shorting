package models

type User struct {
  ID        string  `json:"id"`
  Username  string `json:"username"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Email     string `json:"email"`
  PasswordHash  string `json:"password_hash"`
  CreatedAt int64 `json:"created_at"`
}