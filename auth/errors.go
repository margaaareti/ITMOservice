package auth

import "errors"

var (
	ErrUserFound          = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
)
