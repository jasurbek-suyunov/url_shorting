package postgres

import (
	"context"
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
	randomQrCodePath := uuid.New().String()

	return &models.Url{
		ID:         randomUserID,
		UserID:     genuser().ID,
		OrgPath:    randomOrgPath,
		ShortPath:  randomShortPath,
		Counter:    randomCounter,
		Status:     randomStatus,
		CreatedAt:  int64(randomCreatedAt),
		UpdatedAt:  int64(randomUpdatedAt),
		QrCodePath: randomQrCodePath,
	}
}
func createurl(t *testing.T) (*models.Url, *models.Url, error) {
	mockurl := genurl()
	_, use, _ := createuser(t)
	mockurl.UserID = use.ID
	require.NotEmpty(t, mockurl)

	then, err := strg.Url().CreateUrl(context.Background(), mockurl)
	return mockurl, then, err
}
func TestUrl_CreateUrl(t *testing.T) {
	then, _, err := createurl(t)
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
	require.NotEmpty(t, then.QrCodePath)
}
func TestUrl_GetUrlByID(t *testing.T) {
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
	got, err := strg.Url().GetUrlByID(context.Background(), then.ID)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.NotEmpty(t, got.ID)
	require.NotEmpty(t, got.OrgPath)
	require.NotEmpty(t, got.ShortPath)
	require.NotEmpty(t, got.Counter)
	require.NotEmpty(t, got.CreatedAt)
	require.NotEmpty(t, got.Status)
}
func TestUrl_DeleteUrl(t *testing.T) {
	_, then, _ := createurl(t)

	err := strg.Url().DeleteUrl(context.Background(), then.ShortPath)
	require.NoError(t, err)
}
func TestUrl_UpdateUrl(t *testing.T) {
	_, then, _ := createurl(t)
	then.Counter = 100
	then.Status = 100
	then.UpdatedAt = 100
	then.QrCodePath = "100"
	upd, err := strg.Url().UpdateUrl(context.Background(), then)
	require.NoError(t, err)
	require.NotEmpty(t, upd)
	require.NotEmpty(t, upd.ID)
	require.NotEmpty(t, upd.UserID)
	require.NotEmpty(t, upd.OrgPath)
	require.NotEmpty(t, upd.ShortPath)
	require.NotEmpty(t, upd.Counter)
	require.NotEmpty(t, upd.CreatedAt)
	require.NotEmpty(t, upd.Status)
	require.NotEmpty(t, upd.UpdatedAt)
	require.NotEmpty(t, upd.QrCodePath)
}
