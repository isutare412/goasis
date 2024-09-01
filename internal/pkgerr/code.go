package pkgerr

import (
	"errors"
	"fmt"
	"net/http"
)

type Code int

const (
	CodeUnspecified Code = iota
	CodeNotFound
	CodeInvalidInput
)

func (c Code) HTTPStatusCode() int {
	switch c {
	case CodeUnspecified:
		return http.StatusInternalServerError
	case CodeNotFound:
		return http.StatusNotFound
	case CodeInvalidInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// CodeError wraps original error with custom code.
type CodeError struct {
	// optional error code.
	Code Code

	// original error.
	Err error

	// optional client message.
	ClientMsg string
}

func (e CodeError) Error() string {
	err := fmt.Errorf("error code %d", e.Code)
	if e.ClientMsg != "" {
		err = errors.Join(err, fmt.Errorf("%s", e.ClientMsg))
	}
	if e.Err != nil {
		err = errors.Join(err, e.Err)
	}
	return err.Error()
}

func (e CodeError) Unwrap() error {
	return e.Err
}
