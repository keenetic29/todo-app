package domain

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrTaskNotFound  = errors.New("task not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrUnauthorized = errors.New("unauthorized")
	ErrCategoryNotFound = errors.New("category not found")
    ErrUserAlreadyExists  = errors.New("user with this email already exists")
    ErrUsernameAlreadyTaken = errors.New("username is already taken")
)