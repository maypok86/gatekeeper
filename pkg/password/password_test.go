package password

import (
	"testing"

	"github.com/maypok86/gatekeeper/pkg/random"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := random.String(6)

	firstHashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, firstHashedPassword)

	err = CheckPassword(password, firstHashedPassword)
	require.NoError(t, err)

	wrongPassword := random.String(6)
	err = CheckPassword(wrongPassword, firstHashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	secondHashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, secondHashedPassword)
	require.NotEqual(t, firstHashedPassword, secondHashedPassword)
}
