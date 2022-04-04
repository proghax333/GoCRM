package utils

import (
	"testing"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Before() error {
	godotenv.Load("../.env")
	return nil
}

func TestJwtToken(t *testing.T) {
	err := Before()

	// Check if environment variables are loaded
	if !assert.Equal(t, err, nil, "Error should be nil") {
		return
	}

	data := jwt.MapClaims{
		"userid": 12345,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	}
	token, error := SignToken(data)

	// Check if token creation was successful
	if !assert.Equal(t, error, nil, error) {
		return
	}

	iDecodedToken, error := DecodeToken(token)

	// Check if decoding was successful
	if !assert.Equal(t, error, nil, error) {
		return
	}
	decodedToken, ok := iDecodedToken.(jwt.MapClaims)

	// Check if the decoded data is JSON
	if !ok {
		t.Errorf("Interface to map(json) conversion failed")
		return
	}

	a_userid := data["userid"].(int)
	a_nbf := data["nbf"].(int64)

	b_userid := int(decodedToken["userid"].(float64))
	b_nbf := int64(decodedToken["nbf"].(float64))

	if !(a_userid == b_userid && a_nbf == b_nbf) {
		t.Errorf("Data did not match!")
		return
	}

	t.Logf("JWT Token signed and decoded successfully")
}
