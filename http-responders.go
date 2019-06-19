package core

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
)

// RespondWithError will send a reply with an error as JSON and a HTTP Status code
func RespondWithError(w http.ResponseWriter, code int, err error) {
	props := map[string]string{
		"http_status": strconv.Itoa(code),
		"error": err.Error(),
	}

	var field reflect.Value
	errValue := reflect.ValueOf(err)

	// detect errors like fmt.Errorf()
	if errValue.Type().Kind() == reflect.Ptr {
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
	if field = errValue.FieldByName("Value"); field.IsValid() && !field.IsNil() {
		switch field.Type().Kind() {
		case reflect.String:
			if field.Len() > 0 {
				props["value"] = field.String()
			}
		case reflect.Interface, reflect.Ptr:
			if !field.IsNil() {
				props["value"] = field.String()
			}
		}
	}
	RespondWithJSON(w, code, props)
}

// RespondWithJSON will send a reply with a JSON payload and a HTTP Status code
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
}