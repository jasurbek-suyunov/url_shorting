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

func genurl() *models.Url {

  randomUserID := uuid.New().String()
  randomOrgPath:= uuid.New().String()
  randomShortPath := uuid.New().String()
  randomCounter := rand.Intn(1000)
  randomCreatedAt := rand.Intn(1000)
  randomType := uuid.New().String()


	return &models.Url{
		ID:        randomUserID,
		UserID:    randomUserID,
		OrgPath:   randomOrgPath,
		ShortPath: randomShortPath,
		Counter:   randomCounter,
		CreatedAt: randomCreatedAt,
		Type:      randomType,
	}
}
func createurl(t *testing.T) (*models.Url, *models.Url, error) {
	mockurl 	:= genurl()
	require.NotEmpty(t, mockurl)

	then, err :=strg.Url().CreateUrl(context.Background(), mockurl)
	return mockurl, then, err
}

func TestUrl_CreateUrl(t *testing.T) {
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