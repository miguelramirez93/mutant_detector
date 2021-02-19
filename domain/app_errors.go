package domain

import "errors"

// AppError Single struct that represents errors in logic and data level
type AppError struct {
	Err         error  `json:"error"`
	Description string `json:"description"`
}

// DeliveryError Represents spected error struct for delivery level
type DeliveryError struct {
	Err         string `json:"error"`
	Description string `json:"description"`
}

var (
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("bad-params-input")
)
