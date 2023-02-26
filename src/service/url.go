package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/SuyunovJasurbek/url_shorting/helper"
	"github.com/SuyunovJasurbek/url_shorting/models"
)

func (s *Service) CreateUrl(ctx context.Context, url *models.UrlRequest) (*models.Url, error) {

	var (
		status     int
		short_path string
		user_id    string
		qr_code    string
		expCount   int
		expTime    int64
		expTimeT   time.Time
	)

	user_id = ctx.Value("user_id").(string)

	qr_code, err := helper.GenerateQrCode(url.OrgPath)

	if err != nil {
		log.Printf("Error while generating qr code: %v", err)
		qr_code = ""
	}

	if url.CustomURL != "" {
		short_path = url.CustomURL
	} else {
		short_path = helper.GenerateShortLink(url.OrgPath, user_id)
	}

	if url.ExpCount != "" {
		expCount, err = strconv.Atoi(url.ExpCount)
		if err != nil {
			return nil, errors.New("expCount must be integer")
		} else if expCount <= 0 {
			return nil, errors.New("expCount must be greater than 0")
		}
	}

	if url.ExpTime != "" {
		expTimeT, err = time.Parse("2006-01-02 15:04:05", url.ExpTime)
		if err != nil {
			return nil, errors.New("expTime must be in format 2006-01-02 15:04:05 and must be greater than current time")
		} else if expTimeT.Before(time.Now()) {
			return nil, errors.New("expTime must be greater than current time")
		} else {
			expTime = expTimeT.Unix()
		}
	}

	if url.ExpCount != "" && url.ExpTime != "" {
		url.ExpCount = strconv.Itoa(expCount)
		url.ExpTime = strconv.Itoa(int(expTime))
		status = 3
	} else if url.ExpCount == "" && url.ExpTime != "" {
		status = 2
		url.ExpTime = strconv.Itoa(int(expTime))
	} else if url.ExpCount != "" && url.ExpTime == "" {

		status = 1
		url.ExpCount = strconv.Itoa(expCount)
	} else {
		status = 0
	}

	result, err := s.storage.Url().CreateUrl(ctx, &models.Url{
		UserID:     user_id,
		OrgPath:    url.OrgPath,
		ShortPath:  short_path,
		Counter:    0,
		CreatedAt:  time.Now().Unix(),
		Status:     status,
		QrCodePath: qr_code,
	})

	if err != nil {
		return nil, err
	}

	if result.Status != 0 {
		marshal, err := json.Marshal(url)
		if err != nil {
			return nil, err
		}
		s.cache.Redis().Set(ctx, result.ShortPath, string(marshal), time.Duration(0))
	}

	return result, nil
}

func (s *Service) GetUrl(ctx context.Context, short_path string) (string, error) {

	urlFromRedis, err := s.cache.Redis().Get(ctx, short_path)
	if err == nil {
		var url models.UrlRequest
		err = json.Unmarshal([]byte(urlFromRedis), &url)
		if err != nil {
			log.Println("Unmarshal error: ", err)
			return "", err
		}
		if url.ExpCount != "" {
			expCount, err := strconv.Atoi(url.ExpCount)
			if err != nil {
				log.Println("atoi count error: ", err)
				return "", err
			}
			if expCount <= 0 {
				err := s.cache.Redis().Delete(ctx, short_path)
				if err != nil {
					log.Println("count set error: ", err)
					return "", err
				}
				return "", errors.New("not found")
			}
			expCount--
			url.ExpCount = strconv.Itoa(expCount)
		}
		if url.ExpTime != "" {
			expTime, err := strconv.Atoi(url.ExpTime)
			if err != nil {
				log.Println("atoi time error: ", err)
				return "", err
			}
			if time.Now().Before(time.Unix(int64(expTime), 0)) {
				err := s.cache.Redis().Delete(ctx, short_path)
				if err != nil {
					log.Println("time set error: ", err)
					return "", err
				}
				err = s.storage.Url().DeleteUrlByShortURL(ctx, short_path)
				if err != nil {
					log.Println("delete from repo: ", err)
					return "", err
				}

				return "", errors.New("not found")
			}
		}
		marshal, err := json.Marshal(url)
		if err != nil {
			log.Println("resp unmarshal: ", err)
			return "", err
		}
		s.cache.Redis().Set(ctx, short_path, string(marshal), time.Duration(0))
		return url.OrgPath, nil
	}

	urlFromDb, err := s.storage.Url().GetUrlByShortPath(ctx, short_path)
	if err != nil {
		log.Println("get repo: ", err)
		return "", err
	}
	return urlFromDb.OrgPath, nil
}
func (s *Service) GetUrls(ctx context.Context) ([]*models.Url, error) {
	user_id := ctx.Value("user_id").(string)
	urls, err := s.storage.Url().GetUrls(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return urls, nil
}
func (s *Service) GetUrlByID(ctx context.Context, id string) (*models.Url, error) {
	url, err := s.storage.Url().GetUrlByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return url, nil
}
func (s *Service) DeleteUrl(ctx context.Context, id string) error {
	err := s.storage.Url().DeleteUrlByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) UpdateUrl(ctx context.Context, url *models.Url) (*models.Url, error) {
	return nil, nil
}
