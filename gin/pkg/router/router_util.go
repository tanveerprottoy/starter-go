package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

func GetURLParam(r *http.Request, k string) string {
	return chi.URLParam(r, k)
}

func GetQueryParam(r *http.Request, k string) string {
	return r.URL.Query().Get(k)
}
