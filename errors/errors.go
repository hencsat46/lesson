package custom_errors

import "errors"

var (
	ErrUserExists error = errors.New("user exists")
)
