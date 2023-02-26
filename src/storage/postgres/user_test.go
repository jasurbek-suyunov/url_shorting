package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func createRandomuser(t *testing.T) *models.User {

	user := &models.User{
		Username:     uuid.NewString(),
		FirstName:    uuid.NewString(),
		LastName:     uuid.NewString(),
		Email:        uuid.NewString(),
		PasswordHash: uuid.NewString(),
		CreatedAt:    time.Now().Unix(),
	}

	user2, err := strg.User().CreateUser(context.Background(), user)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, user.Username)
	require.Equal(t, user.CreatedAt, user.CreatedAt)
	return user2
}

func deleteRandomUser(t *testing.T, userID string) {

	err := strg.User().DeleteUser(context.Background(), userID)
	require.NoError(t, err)
}

func Test_CreateUser(t *testing.T) {
	defer deleteRandomUser(t, createRandomuser(t).ID)

}

func Test_DeleteUser(t *testing.T) {

	// bad case 1
	err := strg.User().DeleteUser(context.Background(), "invalid-uuid")
	require.Error(t, err)

	// bad case 2
	err = strg.User().DeleteUser(context.Background(), uuid.NewString())
	require.Error(t, err)

	// good case
	user := createRandomuser(t)
	err = strg.User().DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
}

func Test_GetUserByID(t *testing.T) {

	user := createRandomuser(t)
	defer deleteRandomUser(t, user.ID)

	// bad case 1
	_, err := strg.User().GetUserByID(context.Background(), "invalid-uuid")
	require.Error(t, err)

	// bad case 2
	_, err = strg.User().GetUserByID(context.Background(), uuid.NewString())
	require.Error(t, err)

	// good case
	user2, err := strg.User().GetUserByID(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user.ID, user2.ID)
	require.Equal(t, user.Username, user2.Username)
}

func Test_GetUserByUsername(t *testing.T) {

	user := createRandomuser(t)
	defer deleteRandomUser(t, user.ID)

	// bad case 1
	_, err := strg.User().GetUserByUsername(context.Background(), "invalid-username")
	require.Error(t, err)

	// good case
	user2, err := strg.User().GetUserByUsername(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user.ID, user2.ID)
	require.Equal(t, user.Username, user2.Username)
}

func Test_UpdateUser(t *testing.T) {

	user := createRandomuser(t)
	defer deleteRandomUser(t, user.ID)

	// bad case 1
	user2, err := strg.User().UpdateUser(context.Background(), &models.User{})
	require.Error(t, err)
	require.Empty(t, user2)

	// bad case 2
	user2, err = strg.User().UpdateUser(context.Background(), &models.User{
		ID:        "invalid-uuid",
		Username:  uuid.NewString(),
		FirstName: uuid.NewString(),
	})
	require.Error(t, err)
	require.Empty(t, user2)

	// bad case 3
	user2, err = strg.User().UpdateUser(context.Background(), &models.User{
		ID:        uuid.NewString(),
		Username:  uuid.NewString(),
		FirstName: uuid.NewString(),
	})
	require.Error(t, err)
	require.Empty(t, user2)

	// good case
	user.FirstName = "test"
	user2, err = strg.User().UpdateUser(context.Background(), user)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user.ID, user2.ID)
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.FirstName, user2.FirstName)
}
