package models
type Url struct {
  ID        string `json:"id" db:"id"`
  UserID    string `json:"user_id" db:"user_id"`
  OrgPath   string `json:"org_path" db:"org_path"`
  ShortPath string `json:"short_path" db:"short_path"`
  Counter   int    `json:"counter" db:"counter"`
  CreatedAt int    `json:"created_at" db:"created_at"`
  Type      string `json:"type" db:"type"`
}

type UrlRequest struct {
  UserID  string` json:"user_id" db:"user_id"`
  OrgPath string` json:"org_path" db:"org_path"`
}