package response

import (
	"encoding/json"
	"net/http"
)

func writeResponse(writer http.ResponseWriter, bytes []byte) {
	_, _ = writer.Write(bytes)
}

func BuildData[T any](payload T) *Response[T] {
	return &Response[T]{Data: payload}
}

func Respond(code int, payload any, writer http.ResponseWriter) {
	res, err := json.Marshal(payload)
	if err != nil {
		RespondError(http.StatusInternalServerError, err, writer)
		return
	}
	writer.WriteHeader(code)
	writeResponse(writer, res)
}

func RespondError(code int, err error, writer http.ResponseWriter) {
	res, err := json.Marshal(map[string]string{"error": err.Error()})
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(code)
	if err != nil {
		writeResponse(writer, []byte(err.Error()))
		return
	}
	writeResponse(writer, res)
}

func RespondErrorMessage(code int, msg string, writer http.ResponseWriter) {
	res, err := json.Marshal(map[string]string{"error": msg})
	writer.WriteHeader(code)
	if err != nil {
		writeResponse(writer, []byte(err.Error()))
		return
	}
	writeResponse(writer, res)
}
