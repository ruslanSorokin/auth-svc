package local

import (
	"context"
	"testing"

	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
	"github.com/ruslanSorokin/authentication-service/pkg/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	repo *SessionStore
)

func init() {
	repo = NewSessionStore()
}

func TestBasic(t *testing.T) {
	accountID := "228"
	want := &model.Session{
		AccountID: accountID,
	}

	assert := assert.New(t)
	require := require.New(t)

	id, err := repo.CreateSession(context.Background(), want)
	assert.Nil(err, "Must be nil")

	got, err := repo.GetSessionByID(context.Background(), id)
	assert.Nil(err, "Must be nil")
	require.Equal(*want, *got, "Must be equal")

	got, err = repo.GetSessionByAccountID(context.Background(), accountID)
	assert.Nil(err, "Must be nil")
	require.Equal(*want, *got, "Must be equal")

	err = repo.DeleteSessionByID(context.Background(), id)
	assert.Nil(err, "Must be nil")

	err = repo.DeleteSessionByAccountID(context.Background(), accountID)
	assert.EqualError(err, repository.ErrSessionNotFound.Error(), "Must be Error")
}
