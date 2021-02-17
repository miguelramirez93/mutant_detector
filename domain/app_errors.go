package domain

import "errors"

// AppError Single struct that represents spected errors in the app
type AppError struct {
	Err         error
	Description string
}

var (
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("bad-params-input")
)
