package xero

import "fmt"

type ValidationError struct {
	Message string `json:"Message"`
}

type ErrorElement struct {
	ValidationErrors []ValidationError `json:"ValidationErrors"`
}

type BadRequestError struct {
	ErrorNumber int            `json:"ErrorNumber"`
	Type        string         `json:"Type"`
	Message     string         `json:"Message"`
	Elements    []ErrorElement `json:"Elements,omitempty"`
}

func (e *BadRequestError) Error() string {
	return fmt.Sprintf("request error %v", e.Message)
}
