// Package errors provides error handling utilities with wrapping functionality.
package errors

import (
	"errors"
	"fmt"
)

// New creates a new error with the given message.
func New(message string) error {
	return errors.New(message)
}

func Unwarp(err error) error {
	return errors.Unwrap(err)
}

// Join combines multiple errors into a single error.
func Join(err ...error) error {
	return errors.Join(err...)
}

type simpleerror struct {
	issue string
	cause error
}

func (e *simpleerror) Unwrap() error {
	return e.cause
}

func (s *simpleerror) Error() string {
	if s.cause != nil {
		return fmt.Sprintf("%s; %s", s.issue, s.cause.Error())
	}

	return s.issue
}

func (s *simpleerror) Cause() error {
	return s.cause
}

// Wrap wraps an error with additional context message.
func Wrap(cause error, message string) error {
	return &simpleerror{
		issue: message,
		cause: cause,
	}
}

// Wrapf wraps an error with a formatted context message.
func Wrapf(cause error, format string, args ...any) error {
	return &simpleerror{
		issue: fmt.Sprintf(format, args...),
		cause: cause,
	}
}

// ErrorWithCause defines an interface for errors that have a cause.
type ErrorWithCause interface {
	Error() string
	Cause() error
}

// ErrNotImplemented should be thrown to indicate that a particular
// feature isn't implemented
var ErrNotImplemented = errors.New("not implemented")

// ErrUnsupported is an alias for Go's standard ErrUnsupported error.
var ErrUnsupported = errors.ErrUnsupported
