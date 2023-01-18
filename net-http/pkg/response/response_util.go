package response

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, b []byte) {
	_, _ = w.Write(b)
}

func BuildData[T any](p T) *Response[T] {
	return &Response[T]{
		Data: p,
	}
}

func Respond(c int, p any, w http.ResponseWriter) {
	response, err := json.Marshal(p)
	if err != nil {
		RespondError(http.StatusInternalServerError, err, w)
		return
	}
	w.WriteHeader(c)
	writeResponse(w, response)
}

func RespondError(c int, err error, w http.ResponseWriter) {
	response, err := json.Marshal(map[string]string{"error": err.Error()})
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(c)
	if err != nil {
		writeResponse(w, []byte(err.Error()))
		return
	}
	writeResponse(w, response)
}

func RespondErrorMessage(c int, msg string, w http.ResponseWriter) {
	response, err := json.Marshal(map[string]string{"error": msg})
	w.WriteHeader(c)
	if err != nil {
		writeResponse(w, []byte(err.Error()))
		return
	}
	writeResponse(w, response)
}
