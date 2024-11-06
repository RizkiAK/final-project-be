package helpers

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateUUID() (uint32, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return 0, err
	}
	return id.ID(), nil
}

func GenerateJWTTokenAdmin(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
