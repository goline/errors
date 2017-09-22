package errors

import (
	"fmt"
	e "github.com/pkg/errors"
	"net/http"
)

// Error represents for a common error
type Error interface {
	ErrorMessageAware
	ErrorHttpAware
	ErrorCodeAware
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

type ErrorHttpAware interface {
	// Status returns HTTP Status code
	Status() int

	// WithStatus sets HTTP Status code
	WithStatus(status int) ErrorHttpAware
}

var errStringFormat = "[%s] %s"

func New(code string, message string) Error {
	return &FactoryError{
		code:    code,
		message: message,
		stack:   e.Errorf(errStringFormat, code, message),
		status:  http.StatusOK,
	}
}

type FactoryError struct {
	code    string
	message string
	stack   error
	status  int
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

func (e *FactoryError) Status() int {
	return e.status
}

func (e *FactoryError) WithStatus(status int) ErrorHttpAware {
	e.status = status
	return e
}

// Error implements error interface
func (e *FactoryError) Error() string {
	return fmt.Sprintf(errStringFormat, e.code, e.message)
}
