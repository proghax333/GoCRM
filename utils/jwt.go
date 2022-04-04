package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func GetSecret() (string, error) {
	return GetEnv("JWT_TOKEN_KEY"), nil
}

func SignToken(data jwt.MapClaims) (string, error) {
	secret_string, err := GetSecret()

	if err != nil {
		return "", err
	}

	secret := []byte(secret_string)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func DecodeToken(tokenString string) (interface{}, error) {
	secret, err := GetSecret()

	if err != nil {
		return nil, nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token supplied")
	}
}
