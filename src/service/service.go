package service

import (
	"github.com/SuyunovJasurbek/url_shorting/src/storage"
)

type Service struct {
	storage storage.StorageI
	cache   storage.CacheStorageI
}

func NewService(repo storage.StorageI, redis storage.CacheStorageI) *Service {
	return &Service{
		storage: repo,
		cache:   redis,
	}
}
