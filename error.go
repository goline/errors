package errors

import (
	"fmt"
	"net/http"

	e "github.com/pkg/errors"
)

// Error represents for a common error
type Error interface {
	ErrorMessageAware
	ErrorHttpAware
	ErrorCodeAware
	ErrorDebugger
	ErrorLeveller
	ErrorTracer
	error
}

type ErrorCodeAware interface {
	// Code returns error's code
	Code() string

	// WithCode sets error's code
	WithCode(code string) Error
}

type ErrorMessageAware interface {
	// Message returns error's message
	Message() string

	// WithMessage sets error's message
	WithMessage(message string) Error
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
	WithStatus(status int) Error
}

type ErrorLeveller interface {
	// Level returns error's level
	// it is useful for debugging
	Level() string

	// WithLevel sets error's level
	WithLevel(level string) Error
}

type ErrorDebugger interface {
	// Debug returns message for debugging
	Debug() string

	// WithDebug sets debug's message
	// It is useful to hide system debug message from response
	WithDebug(debug string) Error
}

var errStringFormat = "[%s] %s"

func New(code string, message string) Error {
	return &FactoryError{
		code:    code,
		message: message,
		stack:   e.Errorf(errStringFormat, code, message),
		status:  http.StatusOK,
		level:   LEVEL_ERROR,
	}
}

type FactoryError struct {
	code    string
	message string
	stack   error
	status  int
	level   string
	debug   string
}

func (e *FactoryError) Code() string {
	return e.code
}

func (e *FactoryError) WithCode(code string) Error {
	e.code = code
	return e
}

func (e *FactoryError) Message() string {
	return e.message
}

func (e *FactoryError) WithMessage(message string) Error {
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

func (e *FactoryError) WithStatus(status int) Error {
	e.status = status
	return e
}

func (e *FactoryError) Level() string {
	return e.level
}

func (e *FactoryError) WithLevel(level string) Error {
	e.level = level
	return e
}

// Error implements error interface
func (e *FactoryError) Error() string {
	return fmt.Sprintf(errStringFormat, e.code, e.message)
}

func (e *FactoryError) Debug() string {
	return e.debug
}

func (e *FactoryError) WithDebug(debug string) Error {
	e.debug = debug
	return e
}
