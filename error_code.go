package epayments

import (
	"errors"
	"net/http"
)

// Refer to 2.6
// ErrorCode accepts error exception
type ErrorCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorCode) Validate() (int, error) {
	if e.Code == "0" {
		return http.StatusOK, nil
	}

	err := errors.New(e.Message)
	return http.StatusInternalServerError, err
}
