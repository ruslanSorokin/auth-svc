package local

import (
	"context"
	"testing"

	"github.com/ruslanSorokin/auth-service/pkg/domain/model"
	"github.com/ruslanSorokin/auth-service/pkg/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	repo *SessionRepository
)

func init() {
	repo = NewSessionRepository()
}

func Test(t *testing.T) {
	userID := "228"
	want := &model.Session{
		UserID: userID,
	}

	assert := assert.New(t)
	require := require.New(t)

	id, err := repo.CreateSession(context.Background(), want)
	assert.Nil(err, "Must be nil")

	got, err := repo.GetSessionByID(context.Background(), id)
	assert.Nil(err, "Must be nil")
	require.Equal(*want, *got, "Must be equal")

	got, err = repo.GetSessionByUserID(context.Background(), &userID)
	assert.Nil(err, "Must be nil")
	require.Equal(*want, *got, "Must be equal")

	err = repo.DeleteSessionByID(context.Background(), id)
	assert.Nil(err, "Must be nil")

	err = repo.DeleteSessionByUserID(context.Background(), &userID)
	assert.EqualError(err, repository.ErrSessionNotFound.Error(), "Must be Error")
}
