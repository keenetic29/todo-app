package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTUtil interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(tokenString string) (uint, error)
}

type jwtUtil struct {
	secret string
}

func NewJWTUtil(secret string) JWTUtil {
	return &jwtUtil{secret: secret}
}

func (j *jwtUtil) GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(j.secret))
}

func (j *jwtUtil) ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		return userID, nil
	}

	return 0, errors.New("invalid token")
}