package utils

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Before() error {
	godotenv.Load(".env")
	return nil
}

func TestSignToken(t *testing.T) {
	err := Before()

	if !assert.Equal(t, err, nil, "Error should be nil") {
		return
	}

}
