package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	apiKey := header.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("authorization header is not found")
	}
	apiKey, found := strings.CutPrefix(apiKey, "ApiKey ")
	if !found {
		return "", errors.New("malformed auth header, expecting \"ApiKey <key>\"")
	}
	return apiKey, nil
}
