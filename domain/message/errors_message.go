package message

import (
	"errors"
)

var (
	ErrPageNotFound                = errors.New("404 page not found")
	ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password")
	ErrUnexpectedSigningMethod     = errors.New("unexpected signing method")
	ErrInvalidToken                = errors.New("invalid token")
	ErrInvalidRefreshToken         = errors.New("invalid refresh token")
	ErrInvalidTokenSignature       = errors.New("invalid token signature")
	ErrExpiredToken                = errors.New("expired token")
	ErrCreatingMessage             = errors.New("error creating message")
	ErrItemNotFound                = errors.New("item not found")
	ErrLoadingEnvFile              = errors.New("error loading .env file")
	ErrUsernameIsRequired          = errors.New("username is required")
	ErrPasswordIsRequired          = errors.New("password is required")
	ErrNameIsRequired              = errors.New("name is required")
	ErrEmailIsRequired             = errors.New("email is required")
)
