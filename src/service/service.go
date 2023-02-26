package service

import (
	"github.com/SuyunovJasurbek/url_shorting/src/storage"
)

type Service struct {
	repo storage.StorageI
}

func NewService(repo storage.StorageI) *Service {
	return &Service{
		repo: repo,
	}
}
