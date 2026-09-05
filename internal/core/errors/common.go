package errors

import "errors"

// auth
var ErrInvalidInput = errors.New("invalid input")
var ErrLoginTaken = errors.New("login taken")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInvalidToken = errors.New("invalid token")

// orders
var ErrInvalidNumber = errors.New("invalid number")
var ErrOrderAlreadyLoaded = errors.New("order already loaded")
var ErrOrderTakenByAnother = errors.New("order taken by another")

// accrual
var ErrNoContent = errors.New("no content")
var ErrTooManyRequest = errors.New("too many request")
var ErrInternalServer = errors.New("internal server")
