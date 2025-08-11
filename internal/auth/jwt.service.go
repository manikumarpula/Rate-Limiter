package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secretKey string
	subject   string
}

// NewJwtService creates a new JwtService with the given secret key
func NewJwtService(secretKey string, subject string) JwtService {
	return &jwtService{
		secretKey: secretKey,
		subject:   subject,
	}
}

type JwtService interface {
	DecodeToken(token string) (jwt.MapClaims, error)
}

func (s *jwtService) DecodeToken(tokenString string) (jwt.MapClaims, error) {
	if s.secretKey == "" {
		return "", errors.New("secret key is required")
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
