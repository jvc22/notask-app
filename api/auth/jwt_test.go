package auth

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var testUserId = "testuserid"
var testUsername = "testusername"

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	m.Run()
}

func TestJWTTokenMethods(t *testing.T) {
	t.Run("Invalid token", func(t *testing.T) {
		invalidToken := "invalidtoken"

		_, err := ParseToken(invalidToken)

		assert.NotEqual(t, nil, err)
	})

	t.Run("Wrong comparison", func(t *testing.T) {
		token, _ := GenerateJWT(testUsername)

		parsedTestUserId, _ := ParseToken(token)

		assert.NotEqual(t, testUserId, parsedTestUserId)
	})

	t.Run("Parse token successfully", func(t *testing.T) {
		token, _ := GenerateJWT(testUserId)

		parsedTestUserId, _ := ParseToken(token)

		assert.Equal(t, testUserId, parsedTestUserId)
	})
}
