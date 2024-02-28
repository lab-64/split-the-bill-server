package domain

import "errors"

var ErrNotAuthorized = errors.New("not Authorized")
var InvalidCredentials = errors.New("invalid credentials")
var ErrNotAGroupMember = errors.New("contributor is not a member of the group")
