package service

import (
	"context"
	"time"

	"github.com/SuyunovJasurbek/url_shorting/helper"
	"github.com/SuyunovJasurbek/url_shorting/models"
)

func (s *Service) CreateUser(ctx context.Context, user *models.UserSignUpRequest) (*models.User, error) {

	// generate password hash
	pass_hash, err := helper.GeneratePasswordHash(user.Password)

	// check error
	if err != nil {
		return nil, err
	}

	// create user
	result, err := s.repo.User().CreateUser(ctx, &models.User{
		Username:     user.Username,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PasswordHash: pass_hash,
		CreatedAt:    time.Now().Unix(),
	})

	// check error
	if err != nil {
		return nil, err
	}

	// return result if no error
	return result, nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {

	// get user by username
	result, err := s.repo.User().GetUserByUsername(ctx, username)

	// check error
	if err != nil {
		return nil, err
	}

	// return result if no error
	return result, nil
}
