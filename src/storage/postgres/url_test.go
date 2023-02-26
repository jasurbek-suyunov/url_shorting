package postgres

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func genurl() *models.Url {

	randomUserID := uuid.New().String()
	randomOrgPath := uuid.New().String()
	randomShortPath := uuid.New().String()
	randomCounter := rand.Intn(1000)
	randomCreatedAt := rand.Intn(1000)
	randomUpdatedAt := rand.Intn(1000)
	randomStatus := rand.Intn(1000)

	return &models.Url{
		ID:        randomUserID,
		UserID:    genuser().ID,
		OrgPath:   randomOrgPath,
		ShortPath: randomShortPath,
		Counter:   randomCounter,
		Status:    randomStatus,
		CreatedAt: int64(randomCreatedAt),
		UpdatedAt: int64(randomUpdatedAt),
	}
}
func createurl(t *testing.T) (*models.Url, *models.Url, error) {
	mockurl := genurl()
	_, use, _ := createuser(t)
	mockurl.UserID = use.ID

	fmt.Println(mockurl)
	require.NotEmpty(t, mockurl)

	then, err := strg.Url().CreateUrl(context.Background(), mockurl)
	return mockurl, then, err
}

func TestUrl_CreateUrl(t *testing.T) {
	_, then, err := createurl(t)
	require.NoError(t, err)
	require.NotEmpty(t, then)
	require.NotEmpty(t, then.ID)
	require.NotEmpty(t, then.UserID)
	require.NotEmpty(t, then.OrgPath)
	require.NotEmpty(t, then.ShortPath)
	require.NotEmpty(t, then.Counter)
	require.NotEmpty(t, then.CreatedAt)
	require.NotEmpty(t, then.Status)
	require.NotEmpty(t, then.UpdatedAt)
}

func TestUrl_GetUrlByID(t *testing.T) {
	mockurl, then, err := createurl(t)
	fmt.Println(err.Error())
	require.NoError(t, err)
	require.NotEmpty(t, then)
	require.NotEmpty(t, then.ID)
	require.NotEmpty(t, then.UserID)
	require.NotEmpty(t, then.OrgPath)
	require.NotEmpty(t, then.ShortPath)
	require.NotEmpty(t, then.Counter)
	require.NotEmpty(t, then.CreatedAt)
	require.NotEmpty(t, then.Status)

	got, err := strg.Url().GetUrlByID(context.Background(), mockurl.ID)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.NotEmpty(t, got.ID)
	require.NotEmpty(t, got.UserID)
	require.NotEmpty(t, got.OrgPath)
	require.NotEmpty(t, got.ShortPath)
	require.NotEmpty(t, got.Counter)
	require.NotEmpty(t, got.CreatedAt)
	require.NotEmpty(t, got.Status)
}
