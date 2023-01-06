package db_test

import (
	"context"
	"testing"
	"time"

	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewSession(t *testing.T, ID uuid.UUID) *db.Session {
	arg := &db.NewSessionParams{
		ID:        uuid.New(),
		User:      ID,
		Token:     util.RandomString(10),
		Ip:        util.RandomString(10),
		Agent:     util.RandomString(10),
		ExpiresAt: time.Now(),
	}

	session, err := testQueries.NewSession(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.NotEmpty(t, session.ID)
	require.Equal(t, arg.User, session.User)
	require.Equal(t, arg.Token, session.Token)
	require.Equal(t, arg.Ip, session.Ip)
	require.Equal(t, arg.Agent, session.Agent)
	require.WithinDuration(t, arg.ExpiresAt, session.ExpiresAt, time.Second)
	require.WithinDuration(t, time.Now(), session.CreatedAt, time.Second)

	return session
}

func DeleteSession(t *testing.T, ID uuid.UUID) *db.Session {
	session, err := testQueries.DeleteSession(context.Background(), ID)
	require.NoError(t, err)
	require.NotEmpty(t, session)
	require.Equal(t, ID, session.ID)

	return session
}

func TestNewSessions(t *testing.T) {
	group := NewGroup(t).ID
	session := NewSession(t, NewUser(t, group).ID)
	DeleteSession(t, session.ID)
	DeleteUser(t, session.User)
	DeleteGroup(t, group)
}

func TestGetSession(t *testing.T) {
	group := NewGroup(t).ID
	session := NewSession(t, NewUser(t, group).ID)

	res, err := testQueries.GetSession(context.Background(), session.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, session, res)

	DeleteSession(t, session.ID)
	DeleteUser(t, session.User)
	DeleteGroup(t, group)
}

func TestFindSession(t *testing.T) {
	user := NewUser(t, NewGroup(t).ID)

	for i := 0; i < 10; i++ {
		NewSession(t, user.ID)
	}

	sessions, err := testQueries.FindSession(context.Background(), user.ID)

	require.NoError(t, err)
	require.Len(t, sessions, 10)

	for _, session := range sessions {
		require.NotEmpty(t, session)
		DeleteSession(t, session.ID)
	}

	DeleteUser(t, user.ID)
	DeleteGroup(t, user.Group)
}

func TestListSession(t *testing.T) {
	group := NewGroup(t).ID
	for i := 0; i < 10; i++ {
		NewSession(t, NewUser(t, group).ID)
	}

	list, err := testQueries.ListSessionPage(context.Background(), &db.ListSessionPageParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.Len(t, list, 5)

	for _, session := range list {
		require.NotEmpty(t, session)
	}

	list, err = testQueries.ListSession(context.Background())

	require.NoError(t, err)
	require.Len(t, list, 10)

	for _, session := range list {
		require.NotEmpty(t, session)
		DeleteSession(t, session.ID)
		DeleteUser(t, session.User)
	}

	DeleteGroup(t, group)
}
