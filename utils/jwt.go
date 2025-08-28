package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET")) // Public/exported

type Claims struct {
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token valid for 24 hours
func GenerateJWT(userID int32, username string, secretKey []byte) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ParseJWT validates token and returns claims
func ParseJWT(tokenString string, secretKey []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}
