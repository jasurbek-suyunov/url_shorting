package models

type User struct {
  ID        int64  `json:"id"`
  Username  string `json:"username"`
  Frst_name string `json:"frst_name"`
  Last_name string `json:"last_name"`
  Email     string `json:"email"`
  Password  string `json:"password_hash"`
  Created_at string `json:"created_at"`
}