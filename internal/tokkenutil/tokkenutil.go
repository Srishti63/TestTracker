package tokkenutil

import (
	"fmt"
	"test_tracker_backend/domain"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct{
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func CreateAccessToken(user *domain.User, secret string, erxpiry int )(string, error){
	expirationTime := time.Now().Add(time.Duration(erxpiry)*time.Hour)

	Claims := jwtCustomClaims{
		UserId: user.ID,
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token payload securely: %w", err)
	}

	return tokenString, nil

}
