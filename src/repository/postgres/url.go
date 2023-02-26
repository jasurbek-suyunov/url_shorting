package postgres

import (
	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/jmoiron/sqlx"
)

type urlRepo struct {
	*sqlx.DB
}

// CreateUrl implements repository.UrlI
func (*urlRepo) CreateUrl(user models.Url) (models.Url, error) {
	panic("unimplemented")
}

func NewUrlRepo(db *sqlx.DB) *urlRepo {
	return &urlRepo{db}
}
