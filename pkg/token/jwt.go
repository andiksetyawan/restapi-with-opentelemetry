package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"restapi-with-opentelemetry/config"
)

func JwtTokenIsValid(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SecretToken), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, err
}

func BuildJwtToken(ID, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    ID,
		"exp":   time.Now().UTC().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(config.SecretToken))
}
