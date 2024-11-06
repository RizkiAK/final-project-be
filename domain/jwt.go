package domain

import "github.com/golang-jwt/jwt/v4"

type JWTClaimAdmin struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}
