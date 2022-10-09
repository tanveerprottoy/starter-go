package util

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, b []byte) {
	_, _ = w.Write(b)
}

func Respond(
	c int,
	p interface{},
	w http.ResponseWriter,
) {
	response, err := json.Marshal(p)
	if err != nil {
		RespondError(http.StatusInternalServerError, err, w)
		return
	}
	w.WriteHeader(c)
	writeResponse(w, response)
}

func RespondError(
	c int,
	err error,
	w http.ResponseWriter,
) {
	response, err := json.Marshal(map[string]string{"error": err.Error()})
	w.WriteHeader(c)
	if err != nil {
		writeResponse(w, []byte(err.Error()))
		return
	}
	writeResponse(w, response)
}

func RespondErrorMessage(
	c int,
	msg string,
	w http.ResponseWriter,
) {
	response, err := json.Marshal(map[string]string{"error": msg})
	w.WriteHeader(c)
	if err != nil {
		writeResponse(w, []byte(err.Error()))
		return
	}
	writeResponse(w, response)
}