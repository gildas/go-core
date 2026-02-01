package core

import (
	"encoding/json"
	"html/template"
	"net/http"
	"reflect"
	"strconv"

	"github.com/google/uuid"
)

// RespondWithError will send a reply with an error as JSON and a HTTP Status code
func RespondWithError(w http.ResponseWriter, code int, err error) {
	props := map[string]string{
		"http_status": strconv.Itoa(code),
		"error":       err.Error(),
	}

	var field reflect.Value
	errValue := reflect.ValueOf(err)

	// detect errors like fmt.Errorf()
	if errValue.Type().Kind() == reflect.Pointer {
		errValue = errValue.Elem()
	}

	field = errValue.FieldByName("ID")
	if field.IsValid() && field.Type().Kind().String() == "string" && field.Len() > 0 {
		props["id"] = field.String()
	}
	field = errValue.FieldByName("What")
	if field.IsValid() && field.Type().Kind().String() == "string" && field.Len() > 0 {
		props["what"] = field.String()
	}
	if field = errValue.FieldByName("Value"); field.IsValid() {
		if stringer, ok := field.Interface().(interface{ String() string }); ok {
			props["value"] = stringer.String()
		} else if identifiable, ok := field.Interface().(interface{ GetID() uuid.UUID }); ok {
			props["value"] = identifiable.GetID().String()
		} else if field.Type().Kind().String() == "string" && field.Len() > 0 {
			props["value"] = field.String()
		}
	}
	RespondWithJSON(w, code, props)
}

// RespondWithJSON will send a reply with a JSON payload and a HTTP Status code
//
// The payload will be marshaled using the standard json.Marshal function.
func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

// RespondWithHTMLTemplate will send a reply with a HTML payload generated from an HTML Template and a HTTP Status code
func RespondWithHTMLTemplate(w http.ResponseWriter, code int, template *template.Template, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := template.ExecuteTemplate(w, name, data); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
}
