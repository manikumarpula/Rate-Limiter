package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtService struct {
	secretKey string
	subject   string
}

type RateLimitClaims struct {
	AccountID uuid.UUID `json:"account_id"`
	Subject   string    `json:"sub"`
	Product   string    `json:"product"`
	jwt.RegisteredClaims
}

// NewJwtService creates a new JwtService with the given secret key
func NewJwtService(secretKey string, subject string) JwtService {
	return &jwtService{
		secretKey: secretKey,
		subject:   subject,
	}
}

type JwtService interface {
	DecodeToken(token string) (RateLimitClaims, error)
}

func (s *jwtService) DecodeToken(tokenString string) (RateLimitClaims, error) {
	if s.secretKey == "" {
		return RateLimitClaims{}, errors.New("secret key is required")
	}
	claims := RateLimitClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return RateLimitClaims{}, err
	}
	return claims, nil
}
