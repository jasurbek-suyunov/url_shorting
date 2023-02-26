package repository

import (
	"context"

	"github.com/SuyunovJasurbek/url_shorting/models"
)

type StorageI interface {
	Token() TokenI 
	User()  UserI 
	Url() UrlI
}

type TokenI interface {
}

type UserI interface {
	CreateUser(ctx context.Context,  url *models.User) (*models.User, error)
}

type UrlI interface {
	CreateUrl(ctx context.Context, user *models.Url) (*models.Url, error)
}