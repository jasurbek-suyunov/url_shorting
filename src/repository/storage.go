package repository

import "github.com/SuyunovJasurbek/url_shorting/models"

type StorageI interface {
	Token() TokenI 
	User()  UserI 
	Url() UrlI
}

type TokenI interface {
}

type UserI interface {
	CreateUser(url models.Url) (models.User, error)
}

type UrlI interface {
	CreateUrl(user models.Url) (models.Url, error)
}