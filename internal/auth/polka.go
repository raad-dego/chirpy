package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetPolkaApi(headers http.Header) (string, error) {
	authPolkaApi := headers.Get("Authorization")
	if authPolkaApi == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authPolkaApi, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
