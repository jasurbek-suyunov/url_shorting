package service

import (
	"context"
	"time"

	"github.com/SuyunovJasurbek/url_shorting/helper"
	"github.com/SuyunovJasurbek/url_shorting/models"
)

func (s *Service) CreateUrl(ctx context.Context, url *models.UrlRequest) (*models.Url, error) {

	var (
		status     int
		short_path string
		user_id    string
	)

	user_id = ctx.Value("user_id").(string)

	if url.CustomURL != "" {
		short_path = url.CustomURL
	} else {
		short_path = helper.GenerateShortLink(url.OrgPath, user_id)
	}

	if url.ExpCount != "" && url.ExpTime != "" {
		status = 0
	} else if url.ExpCount == "" && url.ExpTime != "" {
		status = 1
	} else if url.ExpCount != "" && url.ExpTime == "" {
		status = 2
	} else {
		status = 3
	}

	result, err := s.storage.Url().CreateUrl(ctx, &models.Url{
		UserID:    user_id,
		OrgPath:   url.OrgPath,
		ShortPath: short_path,
		Counter:   0,
		CreatedAt: time.Now().Unix(),
		Type:      status,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
