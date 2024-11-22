package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId string) (string, error) {
	claims := JWTClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "notask-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		panic("JWT_SECRET_KEY not set in .env file")
	}

	signedToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(tokenString string) (string, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		panic("JWT_SECRET_KEY not set in .env file")
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			panic("Unexpected signing method")
		}

		return []byte(jwtSecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid or expired token")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	return claims.UserId, nil
}
