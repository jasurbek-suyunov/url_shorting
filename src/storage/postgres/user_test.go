package postgres

import (
	"context"
	"math/rand"
	"testing"
	_ "github.com/lib/pq"
	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func genuser() *models.User {

  randomUsername := uuid.New().String()
  randomFirstName := uuid.New().String()
  randomLastName := uuid.New().String()
  randomPasswordhash := uuid.New().String()
  randomEmail := uuid.New().String()
  randomCreatedAt := rand.Intn(1000)


	return &models.User{     
		Username:  randomUsername,	
		FirstName: randomFirstName,
		LastName:  randomLastName,
		Email:     randomEmail,
		PasswordHash:  randomPasswordhash,
		CreatedAt: int64(randomCreatedAt),
	}
}
func createuser(t *testing.T) (*models.User, *models.User, error) {
	mockuser 	:= genuser()
	require.NotEmpty(t, mockuser)

	then, err :=strg.User().CreateUser(context.Background(), mockuser)
	return mockuser, then, err
}

func TestUser_CreateUser(t *testing.T) {
	_, then, err := createuser(t)
	require.NoError(t, err)
	require.NotEmpty(t, then)
	require.NotEmpty(t, then.ID)
	require.NotEmpty(t, then.Username)
	require.NotEmpty(t, then.FirstName)
	require.NotEmpty(t, then.LastName)
	require.NotEmpty(t, then.Email)
	require.NotEmpty(t, then.PasswordHash)
	require.NotEmpty(t, then.CreatedAt)
}