package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var ErrJsonDecodeError = errors.New("JSON decoding error")

func ValidateJson[T any](body *io.ReadCloser) (*T, error) {
	var data T
	errJsonParse := json.NewDecoder(*body).Decode(&data)
	if errJsonParse != nil {
		return nil, ErrJsonDecodeError
	}

	validate := validator.New()
	errValidator := validate.Struct(data)
	if errValidator != nil {
		return nil, errValidator.(validator.ValidationErrors)
	}

	return &data, nil
}

/*
* Writes to ResponseWriter error messages
* */
func HttpValidateBodyJson[T any](r *http.Request, w http.ResponseWriter) (*T, error) {
	data, ValidationError := ValidateJson[T](&r.Body)

	if ValidationError != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

		message := []byte("Unhandled error: " + ValidationError.Error())

		if errors.Is(ValidationError, ErrJsonDecodeError) {
			message, _ = json.Marshal("Error validating json from body")
		}

		validatorErrors, isOk := ValidationError.(validator.ValidationErrors)
		if isOk {
			message = []byte{}
			for _, err := range validatorErrors {
				currentMessage := fmt.Sprintf("Error: %s;\n\ton[%s]: %s, `%s`\n", err.Error(), err.Field(), err.Value(), err.Tag())
				message = append(message, []byte(currentMessage)...)
			}
		}

		w.Write(message)
		return nil, ValidationError
	}

	return data, nil
}
