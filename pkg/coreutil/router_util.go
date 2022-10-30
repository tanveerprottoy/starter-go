package coreutil

import (
	"net/http"

	"github.com/go-chi/chi"
)

func GetURLParam(k string, r *http.Request) string {
	return chi.URLParam(r, k)
}
