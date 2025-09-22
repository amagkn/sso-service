package entity

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user_already_exists")
	ErrUserNotFound      = errors.New("user_not_found")
	ErrAppNotFound       = errors.New("app_not_found")
	ErrInvalidAppID      = errors.New("invalid_app_id")
)
