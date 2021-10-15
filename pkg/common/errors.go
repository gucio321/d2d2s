package common

import "errors"

// ErrUnexpectedHeader represents an error which occures when section header is incorrect
var ErrUnexpectedHeader = errors.New("unexpected section header signature")
