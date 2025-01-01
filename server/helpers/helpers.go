package helpers

import (
	"encoding/json"
	"net/http"
	"portfolio/models"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var Validate = validator.New()

func ErrJSONResponse(w http.ResponseWriter, optionalMessage string, status int) {
	message := "there's a problem with the response"
	if optionalMessage != "" {
		message = optionalMessage
	}

	response := models.Response{
		Status:  status,
		Message: message,
		Success: false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func JSONResponse(w http.ResponseWriter, opt string, jsonData ...interface{}) {
	m := "success"
	var d interface{}

	if opt != "" {
		m = opt
	}

	if len(jsonData) > 0 {
		d = jsonData[0]
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: m,
		Success: true,
	}

	if d != nil {
		switch reflect.TypeOf(d).Kind() {
		case reflect.Slice:
			if reflect.ValueOf(d).Len() > 0 {
				response.Data = d
			} else {
				response.Data = []interface{}{}
			}
		default:
			response.Data = d
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewUUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}

func BindValidateJSON(w http.ResponseWriter, r *http.Request, body interface{}) error {
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		ErrJSONResponse(w, err.Error(), http.StatusBadRequest)
		return err
	}

	if err := Validate.Struct(body); err != nil {
		ErrJSONResponse(w, err.Error(), http.StatusBadRequest)
		return err
	}

	return nil
}
