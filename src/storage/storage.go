package storage

import (
	"context"

	"github.com/SuyunovJasurbek/url_shorting/models"
)

type StorageI interface {
	Token() TokenI
	User() UserI
	Url() UrlI
}

type TokenI interface {
}

type UserI interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}

type UrlI interface {
	CreateUrl(ctx context.Context, url *models.Url) (*models.Url, error)
	DeleteUrl(ctx context.Context, url *models.Url) (*models.Url, error)
	GetUrlByID(ctx context.Context, urlID string) (*models.Url, error)
	GetUrls(ctx context.Context, userID string) (*models.GetAllUrl, error)
	UpdateUrl(ctx context.Context, url *models.Url) (*models.Url, error)
}
