package epayments

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorCode) Validate() (int, error) {
	fmt.Println("--------------------", e)
	if e.Code == "0" {
		return http.StatusOK, nil
	}

	err := errors.New(e.Message)
	return http.StatusInternalServerError, err
}
