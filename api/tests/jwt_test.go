package tests

import (
	"notask-app/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWTTokenMethods(t *testing.T) {
	t.Run("Invalid token", func(t *testing.T) {
		invalidToken := "invalidtoken"

		_, err := auth.ParseToken(invalidToken)

		assert.NotEqual(t, nil, err)
	})

	t.Run("Wrong comparison", func(t *testing.T) {
		token, _ := auth.GenerateJWT(testUsername)

		parsedTestUserId, _ := auth.ParseToken(token)

		assert.NotEqual(t, testUserId, parsedTestUserId)
	})

	t.Run("Parse token successfully", func(t *testing.T) {
		token, _ := auth.GenerateJWT(testUserId)

		parsedTestUserId, _ := auth.ParseToken(token)

		assert.Equal(t, testUserId, parsedTestUserId)
	})
}
