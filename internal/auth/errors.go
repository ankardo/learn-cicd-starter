package auth

import (
	"errors"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")
var ErrMalformedAuthorizationHeader = errors.New("malformed authorization header")

