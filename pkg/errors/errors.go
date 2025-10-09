package errors

import "fmt"

type BusinessError struct {
	Code    string
	Message string
	Cause   error
}

func (e *BusinessError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func NewBusinessError(code, message string, cause error) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

func WrapError(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}
