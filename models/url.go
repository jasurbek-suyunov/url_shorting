package models

type Url struct {
	ID        string `json:"id" db:"id"`
	UserID    string `json:"user_id" db:"user_id"`
	OrgPath   string `json:"org_path" db:"org_path"`
	ShortPath string `json:"short_path" db:"short_path"`
	Counter   int    `json:"counter" db:"counter"`
	Status    int `json:"status" db:"status"`
	CreatedAt int64    `json:"created_at" db:"created_at"`
	UpdatedAt int64    `json:"updated_at" db:"updated_at"`
}

type UrlRequest struct {
	UserID  string `json:"user_id" db:"user_id"`
	OrgPath string `json:"org_path" db:"org_path"`
}

type GetAllUrlsResponse struct {
	Data  []*Url `json:"data"`
	Error *Error `json:"error"`
	Code  int    `json:"code"`
	Meta  *Meta  `json:"meta"`
}

type GetAllUrl struct {
	Urls []*Url `json:"urls"`
	Meta *Meta  `json:"meta"`
}
