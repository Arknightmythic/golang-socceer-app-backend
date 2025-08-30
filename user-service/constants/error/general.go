package error

import "errors"



var(
	ErrInternalServerError = errors.New("internal server error")
	ErrSQLError = errors.New("database server failed to execute query")
	ErrTooManyRequests = errors.New("too many requests, please try again later")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidToken = errors.New("invalid token")
	ErrForbidden = errors.New("forbidden, you don't have access to this resource")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}