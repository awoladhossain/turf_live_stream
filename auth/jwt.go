package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// token verification - nestjs same secret and algorithm

// func VerifyToken(tokenString string) (*models.AuthUser, error) {
// 	token, err := jwt.ParseWithClaims()
// }
