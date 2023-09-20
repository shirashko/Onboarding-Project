package error

import (
	"github.com/pkg/errors"
)

type UserNotFoundError struct {
	err error
}

// NewNotFoundError is a factory function used to create instances of the UserNotFoundFound error type.
// The function wraps the NotFoundError instance with a stack trace. This helps capture the stack trace when the
// error is created, which is useful for debugging.
func NewNotFoundError(error error) error {
	return errors.WithStack(UserNotFoundError{err: error})
}

// for UserNotFoundError to implement the Error interface, it should implement method called Error
func (e UserNotFoundError) Error() string {
	return e.err.Error()
}
