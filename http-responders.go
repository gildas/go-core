package core

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// RespondWithError will send a reply with an error as JSON and a HTTP Status code
func RespondWithError(w http.ResponseWriter, code int, errorMessage string) {
        RespondWithJSON(w, code, map[string]string{
                "http_status": strconv.Itoa(code),
                "error": errorMessage,
                "message": errorMessage,
        })
}

// RespondWithJSON will send a reply with a JSON payload and a HTTP Status code
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
}