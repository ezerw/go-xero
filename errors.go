package xero

import "fmt"

// ValidationError represents a list of errors returned by Xero.
type ValidationError struct {
	Message string `json:"Message"`
}

type Warning struct {
	Message string `json:"Message"`
}

// ErrorElement represent one element that failed included in the error response.
type ErrorElement struct {
	ValidationErrors []ValidationError `json:"ValidationErrors"`
	Warnings         []Warning         `json:"Warnings"`
}

// BadRequestError represents the error structure returned by Xero.
type BadRequestError struct {
	ErrorNumber int            `json:"ErrorNumber"`
	Type        string         `json:"Type"`
	Message     string         `json:"Message"`
	Elements    []ErrorElement `json:"Elements,omitempty"`
}

// ResourceNotFoundError represents an error returned to the client when the required resource was not found.
type ResourceNotFoundError struct {
	ID       string `json:"id,omitempty"`
	Resource string `json:"resource"`
}

// ValidationNotEmptyError represents an error returned when a required field is empty.
type ValidationNotEmptyError struct {
	Field string `json:"field"`
}

// InvalidTokenError represents an error returned by Xero when the token is invalid or expired.
type InvalidTokenError struct {
	Title  string `json:"Title"`
	Status int    `json:"Status"`
	Detail string `json:"Detail"`
}

func (e BadRequestError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

func (e ResourceNotFoundError) Error() string {
	return fmt.Sprintf("%s ID: %s not found.", e.Resource, e.ID)
}

func (e ValidationNotEmptyError) Error() string {
	return fmt.Sprintf("field '%s' cannot be empty.", e.Field)
}

func (e InvalidTokenError) Error() string {
	return fmt.Sprintf("xero responded with '%d %s' %s.", e.Status, e.Title, e.Detail)
}
