package mediaos

import "fmt"

// APIError reports a non-200 response by the API
type APIError struct {
	StatusCode int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Server Error Response: %d", e.StatusCode)
}

// Error creates a new APIError instance
func Error(code int) error {
	return &APIError{StatusCode: code}
}
