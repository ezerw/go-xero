package xero

import "fmt"

// ValidationError represents a list of errors returned by Xero.
type ValidationError struct {
	Message string `json:"Message"`
}

// ErrorElement represent one element that failed included in the error response.
type ErrorElement struct {
	ValidationErrors []ValidationError `json:"ValidationErrors"`
}

// BadRequestError represents the error structure returned by Xero.
type BadRequestError struct {
	ErrorNumber int            `json:"ErrorNumber"`
	Type        string         `json:"Type"`
	Message     string         `json:"Message"`
	Elements    []ErrorElement `json:"Elements,omitempty"`
}

func (e BadRequestError) Error() string {
	return fmt.Sprintf("bad request error. %s: %s", e.Type, e.Message)
}
