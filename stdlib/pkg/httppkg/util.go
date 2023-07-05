package httppkg

import (
	"errors"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
)

func GetURLParam(req *http.Request, key string) string {
	return chi.URLParam(req, key)
}

func GetQueryParam(req *http.Request, key string) string {
	return req.URL.Query().Get(key)
}

func ParseAuthToken(r *http.Request) ([]string, error) {
	tkHeader := r.Header.Get("Authorization")
	if tkHeader == "" {
		// Token is missing
		return nil, errors.New("auth token is missing")
	}
	splits := strings.Split(tkHeader, " ")
	// token format is `Bearer {tokenBody}`
	if len(splits) != 2 {
		return nil, errors.New("token format is invalid")
	}
	return splits, nil
}

func GetFile(r *http.Request, k string) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile(k)
}

func BuildURL(base, path string, queriesMap map[string]string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	if path != "" {
		// Path param
		u.Path += path
	}
	if queriesMap != nil {
		// Query params
		p := url.Values{}
		for k, v := range queriesMap {
			p.Add(k, v)
		}
		u.RawQuery = p.Encode()
	}
	return u.String(), nil
}
