package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken_Success(t *testing.T) {
	token, err := GenerateToken(1, "admin")

	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestValidateToken(t *testing.T) {
	token, err := GenerateToken(1, "admin")
	require.NotNil(t, token)
	require.Nil(t, err)

	id, role, err := ValidateToken(token)

	assert.NotNil(t, id)
	assert.NotNil(t, role)
	assert.Nil(t, err)

	assert.Equal(t, 1, id)
	assert.Equal(t, "admin", role)
}
