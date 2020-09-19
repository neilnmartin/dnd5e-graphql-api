package repository

import "fmt"

type invalidIDError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e invalidIDError) Error() string {
	return fmt.Sprintf("Error: [%s]: %s", e.Code, e.Message)
}
