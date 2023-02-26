package models

type Url struct {
	ID        string `json:"id" db:"id"`
	UserID    string `json:"user_id" db:"user_id"`
	OrgPath   string `json:"org_path" db:"org_path"`
	ShortPath string `json:"short_path" db:"short_path"`
	Counter   int    `json:"counter" db:"counter"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	Type      int    `json:"type" db:"type"`
}

type UrlRequest struct {
	OrgPath   string `json:"org_path" binding:"required"`
	ExpTime   string `json:"exp_time"`
	ExpCount  string `json:"exp_count"`
	CustomURL string `json:"custom_url"`
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
