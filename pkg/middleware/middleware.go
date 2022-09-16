package middleware

import "net/http"

// JSONContentTypeMiddleWare content type json setter middleware
func JSONContentTypeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
