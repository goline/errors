package errors

import (
	"fmt"
	e "github.com/pkg/errors"
)

// Error represents for a common error
type Error interface {
	ErrorCodeAware
	ErrorMessageAware
	ErrorTracer
	error
}

type ErrorCodeAware interface {
	// Code returns error's code
	Code() string

	// WithCode sets error's code
	WithCode(code string) ErrorCodeAware
}

type ErrorMessageAware interface {
	// Message returns error's message
	Message() string

	// WithMessage sets error's message
	WithMessage(message string) ErrorMessageAware
}

type ErrorTracer interface {
	// Trace prints error's stack trace
	Trace()

	// TraceString returns error's stack trace as string
	TraceString() string
}

var errStringFormat = "[%s] %s"

func NewError(code string, message string) Error {
	return &FactoryError{code, message, e.Errorf(errStringFormat, code, message)}
}

type FactoryError struct {
	code    string
	message string
	stack   error
}

func (e *FactoryError) Code() string {
	return e.code
}

func (e *FactoryError) WithCode(code string) ErrorCodeAware {
	e.code = code
	return e
}

func (e *FactoryError) Message() string {
	return e.message
}

func (e *FactoryError) WithMessage(message string) ErrorMessageAware {
	e.message = message
	return e
}

func (e *FactoryError) Trace() {
	fmt.Println(e.TraceString())
}

func (e *FactoryError) TraceString() string {
	return fmt.Sprintf("%+v", e.stack)
}

// Error implements error interface
func (e *FactoryError) Error() string {
	return fmt.Sprintf(errStringFormat, e.code, e.message)
}
