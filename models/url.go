package models

type Url struct {
  ID        int64  `json:"id"`
  User_id   int64  `json:"user_id"`
  Org_path   string `json:"org_path"`
  Short_path string `json:"short_path"`
}
