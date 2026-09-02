package errors

import "errors"

var ErrInvalidInput = errors.New("invalid input")
var ErrLoginTaken = errors.New("login taken")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInvalidToken = errors.New("invalid token")
