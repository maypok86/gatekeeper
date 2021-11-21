package password

import (
	"testing"

	"github.com/maypok86/gatekeeper/pkg/random"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password, err := random.String(6)
	require.NoError(t, err)

	firstHashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, firstHashedPassword)

	err = CheckPassword(password, firstHashedPassword)
	require.NoError(t, err)

	wrongPassword, err := random.String(6)
	require.NoError(t, err)
	err = CheckPassword(wrongPassword, firstHashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	secondHashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, secondHashedPassword)
	require.NotEqual(t, firstHashedPassword, secondHashedPassword)
}
