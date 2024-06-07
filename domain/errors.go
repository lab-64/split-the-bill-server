package domain

import "errors"

var ErrNotAuthorized = errors.New("not Authorized")
var InvalidCredentials = errors.New("invalid credentials")
var ErrConcurrentModification = errors.New("concurrent modification")
