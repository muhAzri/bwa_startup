package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta  Meta        `json:"meta"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, data interface{}, errMessage interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	var jsonResponse Response

	if errMessage != "" {
		jsonResponse = Response{
			Meta:  meta,
			Error: errMessage, // Set the error message
		}
	} else {
		jsonResponse = Response{
			Meta: meta,
			Data: data,
		}
	}

	return jsonResponse
}

func FormatError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
