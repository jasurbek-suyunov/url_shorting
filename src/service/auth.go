package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/SuyunovJasurbek/url_shorting/helper"
	"github.com/SuyunovJasurbek/url_shorting/models"
)

func (s *Service) CreateUser(ctx context.Context, user *models.UserSignUpRequest) error {

	// generate password hash
	pass_hash, err := helper.GeneratePasswordHash(user.Password)

	// check error
	if err != nil {
		return err
	}

	// create user
	result, err := s.storage.User().CreateUser(ctx, &models.User{
		Username:     user.Username,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PasswordHash: pass_hash,
		CreatedAt:    time.Now().Unix(),
	})

	// check error
	if err != nil {
		return err
	}
	// convert user to json
	user_string, err := json.Marshal(result)

	// check error
	if err != nil {
		return err
	}

	// set username and user
	err = s.cache.Redis().Set(ctx, result.Username, string(user_string), 0)

	// check error
	if err != nil {
		return err
	}

	// return result if no error
	return nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {

	// result
	var user *models.User

	// get user from redis
	user_string, err := s.cache.Redis().Get(ctx, username)

	// check error
	if err != nil {

		// get user by username from database
		user, err = s.storage.User().GetUserByUsername(ctx, username)

		// check error
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("key: ", user_string)
		err = json.Unmarshal([]byte(user_string), &user)
		if err != nil {
			return nil, err
		}
	}

	// return result if no error
	return user, nil
}

