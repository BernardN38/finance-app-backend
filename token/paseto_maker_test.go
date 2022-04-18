package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPestomaker(t *testing.T) {
	maker, err := NewPasetoMaker("lodrtnufepalofeiwnctvuqmifeasklo")
	require.NoError(t, err)

	username := "eris"
	duration := time.Minute

	issued_at := time.Now()
	expired_at := issued_at.Add(duration)

	token, err := maker.CreateToken(username, 1, duration, false)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload)
	require.Equal(t, username, payload.Username)

	require.WithinDuration(t, issued_at, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expired_at, payload.ExpiredAt, time.Second)

}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker("lodrtnufepalofeiwnctvuqmifeasklo")
	require.NoError(t, err)

	token, err := maker.CreateToken("eris123", 1, -time.Minute, false)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
