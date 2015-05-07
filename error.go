package mediaos

import "fmt"

type APIError struct {
	StatusCode int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Server Error Response: %d", e.StatusCode)
}

func Error(code int) error {
	return &APIError{StatusCode: code}
}
