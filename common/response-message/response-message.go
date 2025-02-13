package response_message

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ResponsePaginate struct {
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

func GetResponse(statusCode int, success bool, message string, data interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Success:    success,
		Message:    message,
		Data:       data,
	}
}

func BindRequestErrorChecking(bindError error) []string {
	errorMessages := []string{}
	var validatorErr validator.ValidationErrors

	if errors.As(bindError, &validatorErr) {
		for _, e := range bindError.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
	} else {
		errorMessages = append(errorMessages, bindError.Error())
	}

	return errorMessages
}
