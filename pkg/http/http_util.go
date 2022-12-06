package http

import (
	"errors"
	"mime/multipart"
	"net/http"
	"strings"
)

func ParseAuthToken(r *http.Request) ([]string, error) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		// Token is missing
		return nil, errors.New("auth token is missing")
	}
	splits := strings.Split(tokenHeader, " ")
	// token format is `Bearer {tokenBody}`
	if len(splits) != 2 {
		return nil, errors.New("token format is invalid")
	}
	return splits, nil
}

func GetFile(r *http.Request, k string) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile(k)
}
