package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

func GetURLParam(req *http.Request, key string) string {
	return chi.URLParam(req, key)
}

func GetQueryParam(req *http.Request, key string) string {
	return req.URL.Query().Get(key)
}
