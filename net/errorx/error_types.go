// Code Autogenerated by Jarvis Makefile. DO NOT EDIT
// - source:   templates/error_types.go.template

package errorx

import (
	"fmt"
	"github.com/jucardi/go-streams/streams"
	"net/http"
	"strings"
)

var (
	// Functions that identify an error to be of a specific error type based on the error message.
	errorTypeIdentifiers = map[int]func(string) bool{
		http.StatusBadRequest:            isBadRequest,
		http.StatusUnauthorized:          isUnauthorized,
		http.StatusForbidden:             isForbidden,
		http.StatusNotFound:              isNotFound,
		http.StatusMethodNotAllowed:      isMethodNotAllowed,
		http.StatusNotAcceptable:         isNotAcceptable,
		http.StatusRequestTimeout:        isRequestTimeout,
		http.StatusConflict:              isConflict,
		http.StatusPreconditionFailed:    isPreconditionFailed,
		http.StatusRequestEntityTooLarge: isRequestEntityTooLarge,
		http.StatusInternalServerError:   isUnhandled,
		http.StatusNotImplemented:        isNotImplemented,
		http.StatusBadGateway:            isBadGateway,
	}
)

// IsErrorType verifies if the provided error matches the provided error type
func IsErrorType(code int, err error) bool {
	if e := convert(err); e == nil {
		return false
	} else {
		return e.Code == int32(code)
	}
}

func convert(err error) *Error {
	ret, _ := err.(*Error)
	return ret
}

// ******************************************************
//    400 Bad Request
// ******************************************************

var (
	badRequestAliases    = []string{"bad_request", "bad request", "invalid"}
	badRequestValidators = []func(string) bool{
		func(s string) bool { return streams.From(badRequestAliases).Contains(strings.ToLower(s)) },
	}
)

// NewBadRequest creates a new `Error` and sets the error code to 400 Bad Request (ErrorType_BAD_REQUEST) and title to the default for that status.
func NewBadRequest(msg string, errs ...error) *Error {
	return newError(1, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), msg, errs...)
}

// WrapBadRequest wraps an error and sets the error code to 400 Bad Request (ErrorType_BAD_REQUEST) and title to the default for that status.
func WrapBadRequest(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusBadRequest, msg...)
}

// IsBadRequest indicates whether the error points to a bad request
func IsBadRequest(err error) bool {
	return IsErrorType(http.StatusBadRequest, err)
}

// WrapBadRequestf wraps an error and sets the error code to 400 Bad Request (ErrorType_BAD_REQUEST) using as message the string format result
// of `format` and `args
func WrapBadRequestf(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusBadRequest, fmt.Sprintf(format, args...))
}

func isBadRequest(msg string) bool {
	return streams.From(badRequestValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    401 Unauthorized
// ******************************************************

var (
	unauthorizedAliases    = []string{"unauthorized", "unauthorized"}
	unauthorizedValidators = []func(string) bool{
		func(s string) bool { return streams.From(unauthorizedAliases).Contains(strings.ToLower(s)) },
	}
)

// NewUnauthorized creates a new `Error` and sets the error code to 401 Unauthorized (ErrorType_UNAUTHORIZED) and title to the default for that status.
func NewUnauthorized(msg string, errs ...error) *Error {
	return newError(1, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), msg, errs...)
}

// WrapUnauthorized wraps an error and sets the error code to 401 Unauthorized (ErrorType_UNAUTHORIZED) and title to the default for that status.
func WrapUnauthorized(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusUnauthorized, msg...)
}

// WrapUnauthorizedf wraps an error and sets the error code to 401 Unauthorized (ErrorType_UNAUTHORIZED) using as message the string format result
// of `format` and `args
func WrapUnauthorizedf(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusUnauthorized, fmt.Sprintf(format, args...))
}

// IsUnauthorized indicates whether the error points to an unauthorized access
func IsUnauthorized(err error) bool {
	return IsErrorType(http.StatusUnauthorized, err)
}

func isUnauthorized(msg string) bool {
	return streams.From(unauthorizedValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    403 Forbidden
// ******************************************************

var (
	forbiddenAliases    = []string{"forbidden", "forbidden", "not allowed"}
	forbiddenValidators = []func(string) bool{
		func(s string) bool { return streams.From(forbiddenAliases).Contains(strings.ToLower(s)) },
	}
)

// NewForbidden creates a new `Error` and sets the error code to 403 Forbidden (ErrorType_FORBIDDEN) and title to the default for that status.
func NewForbidden(msg string, errs ...error) *Error {
	return newError(1, http.StatusForbidden, http.StatusText(http.StatusForbidden), msg, errs...)
}

// WrapForbidden wraps an error and sets the error code to 403 Forbidden (ErrorType_FORBIDDEN) and title to the default for that status.
func WrapForbidden(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusForbidden, msg...)
}

// WrapForbiddenf wraps an error and sets the error code to 403 Forbidden (ErrorType_FORBIDDEN) using as message the string format result
// of `format` and `args
func WrapForbiddenf(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusForbidden, fmt.Sprintf(format, args...))
}

// IsForbidden indicates whether the error points to forbidden access to a resource
func IsForbidden(err error) bool {
	return IsErrorType(http.StatusForbidden, err)
}

func isForbidden(msg string) bool {
	return streams.From(forbiddenValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    404 Not Found
// ******************************************************

var (
	notFoundAliases    = []string{"not_found", "not found", "record not found"}
	notFoundValidators = []func(string) bool{
		func(s string) bool { return streams.From(notFoundAliases).Contains(strings.ToLower(s)) },
	}
)

// NewNotFound creates a new `Error` and sets the error code to 404 Not Found (ErrorType_NOT_FOUND) and title to the default for that status.
func NewNotFound(msg string, errs ...error) *Error {
	return newError(1, http.StatusNotFound, http.StatusText(http.StatusNotFound), msg, errs...)
}

// WrapNotFound wraps an error and sets the error code to 404 Not Found (ErrorType_NOT_FOUND) and title to the default for that status.
func WrapNotFound(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusNotFound, msg...)
}

// WrapNotFoundf wraps an error and sets the error code to 404 Not Found (ErrorType_NOT_FOUND) using as message the string format result
// of `format` and `args
func WrapNotFoundf(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusNotFound, fmt.Sprintf(format, args...))
}

// IsNotFound indicates whether the error points to resource not found
func IsNotFound(err error) bool {
	return IsErrorType(http.StatusNotFound, err)
}

func isNotFound(msg string) bool {
	return streams.From(notFoundValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    405 Method Not Allowed
// ******************************************************

var (
	methodNotAllowedAliases    = []string{"method_not_allowed", "method not allowed"}
	methodNotAllowedValidators = []func(string) bool{
		func(s string) bool { return streams.From(methodNotAllowedAliases).Contains(strings.ToLower(s)) },
	}
)

// NewMethodNotAllowed creates a new `Error` and sets the error code to 405 Method Not Allowed (ErrorType_METHOD_NOT_ALLOWED) and title to the default for that status.
func NewMethodNotAllowed(msg string, errs ...error) *Error {
	return newError(1, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), msg, errs...)
}

// WrapMethodNotAllowed wraps an error and sets the error code to 405 Method Not Allowed (ErrorType_METHOD_NOT_ALLOWED) and title to the default for that status.
func WrapMethodNotAllowed(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusMethodNotAllowed, msg...)
}

// WrapMethodNotAllowedf wraps an error and sets the error code to 405 Method Not Allowed (ErrorType_METHOD_NOT_ALLOWED) using as message the string format result
// of `format` and `args
func WrapMethodNotAllowedf(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusMethodNotAllowed, fmt.Sprintf(format, args...))
}

// IsMethodNotAllowed indicates whether the error points to a method not allowed error
func IsMethodNotAllowed(err error) bool {
	return IsErrorType(http.StatusMethodNotAllowed, err)
}

func isMethodNotAllowed(msg string) bool {
	return streams.From(methodNotAllowedValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    406 Not Acceptable
// ******************************************************

var (
	notAcceptableAliases    = []string{"not_acceptable", "not acceptable", "unacceptable"}
	notAcceptableValidators = []func(string) bool{
		func(s string) bool { return streams.From(notAcceptableAliases).Contains(strings.ToLower(s)) },
	}
)

// NewNotAcceptable creates a new `Error` and sets the error code to 406 Not Acceptable (ErrorType_NOT_ACCEPTABLE) and title to the default for that status.
func NewNotAcceptable(msg string, errs ...error) *Error {
	return newError(1, http.StatusNotAcceptable, http.StatusText(http.StatusNotAcceptable), msg, errs...)
}

// WrapNotAcceptable wraps an error and sets the error code to 406 Not Acceptable (ErrorType_NOT_ACCEPTABLE) and title to the default for that status.
func WrapNotAcceptable(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusNotAcceptable, msg...)
}

// WrapNotAcceptablef wraps an error and sets the error code to 406 Not Acceptable (ErrorType_NOT_ACCEPTABLE) using as message the string format result
// of `format` and `args
func WrapNotAcceptablef(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusNotAcceptable, fmt.Sprintf(format, args...))
}

// IsNotAcceptable indicates whether the error points to a not acceptable error
func IsNotAcceptable(err error) bool {
	return IsErrorType(http.StatusNotAcceptable, err)
}

func isNotAcceptable(msg string) bool {
	return streams.From(notAcceptableValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    408 Request Timeout
// ******************************************************

var (
	requestTimeoutAliases    = []string{"request_timeout", "request timeout"}
	requestTimeoutValidators = []func(string) bool{
		func(s string) bool { return streams.From(requestTimeoutAliases).Contains(strings.ToLower(s)) },
	}
)

// NewRequestTimeout creates a new `Error` and sets the error code to 408 Request Timeout (ErrorType_REQUEST_TIMEOUT) and title to the default for that status.
func NewRequestTimeout(msg string, errs ...error) *Error {
	return newError(1, http.StatusRequestTimeout, http.StatusText(http.StatusRequestTimeout), msg, errs...)
}

// WrapRequestTimeout wraps an error and sets the error code to 408 Request Timeout (ErrorType_REQUEST_TIMEOUT) and title to the default for that status.
func WrapRequestTimeout(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusRequestTimeout, msg...)
}

// WrapRequestTimeoutf wraps an error and sets the error code to 408 Request Timeout (ErrorType_REQUEST_TIMEOUT) using as message the string format result
// of `format` and `args
func WrapRequestTimeoutf(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusRequestTimeout, fmt.Sprintf(format, args...))
}

// IsRequestTimeout indicates whether the error points to a request timeout
func IsRequestTimeout(err error) bool {
	return IsErrorType(http.StatusRequestTimeout, err)
}

func isRequestTimeout(msg string) bool {
	return streams.From(requestTimeoutValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    409 Conflict
// ******************************************************

var (
	conflictAliases    = []string{"conflict", "conflict"}
	conflictValidators = []func(string) bool{
		func(s string) bool { return streams.From(conflictAliases).Contains(strings.ToLower(s)) },
		func(s string) bool { return strings.HasPrefix(s, "Error 1062: Duplicate entry") },
	}
)

// NewConflict creates a new `Error` and sets the error code to 409 Conflict (ErrorType_CONFLICT) and title to the default for that status.
func NewConflict(msg string, errs ...error) *Error {
	return newError(1, http.StatusConflict, http.StatusText(http.StatusConflict), msg, errs...)
}

// WrapConflict wraps an error and sets the error code to 409 Conflict (ErrorType_CONFLICT) and title to the default for that status.
func WrapConflict(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusConflict, msg...)
}

// WrapConflictf wraps an error and sets the error code to 409 Conflict (ErrorType_CONFLICT) using as message the string format result
// of `format` and `args
func WrapConflictf(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusConflict, fmt.Sprintf(format, args...))
}

// IsConflict indicates whether the error points to a resource conflict
func IsConflict(err error) bool {
	return IsErrorType(http.StatusConflict, err)
}

func isConflict(msg string) bool {
	return streams.From(conflictValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    412 Precondition Failed
// ******************************************************

var (
	preconditionFailedAliases    = []string{"precondition_failed", "precondition failed"}
	preconditionFailedValidators = []func(string) bool{
		func(s string) bool { return streams.From(preconditionFailedAliases).Contains(strings.ToLower(s)) },
	}
)

// NewPreconditionFailed creates a new `Error` and sets the error code to 412 Precondition Failed (ErrorType_PRECONDITION_FAILED) and title to the default for that status.
func NewPreconditionFailed(msg string, errs ...error) *Error {
	return newError(1, http.StatusPreconditionFailed, http.StatusText(http.StatusPreconditionFailed), msg, errs...)
}

// WrapPreconditionFailed wraps an error and sets the error code to 412 Precondition Failed (ErrorType_PRECONDITION_FAILED) and title to the default for that status.
func WrapPreconditionFailed(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusPreconditionFailed, msg...)
}

// WrapPreconditionFailedf wraps an error and sets the error code to 412 Precondition Failed (ErrorType_PRECONDITION_FAILED) using as message the string format result
// of `format` and `args
func WrapPreconditionFailedf(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusPreconditionFailed, fmt.Sprintf(format, args...))
}

// IsPreconditionFailed indicates whether the error points to a precondition failure
func IsPreconditionFailed(err error) bool {
	return IsErrorType(http.StatusPreconditionFailed, err)
}

func isPreconditionFailed(msg string) bool {
	return streams.From(preconditionFailedValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    413 Request Entity Too Large
// ******************************************************

var (
	requestEntityTooLargeAliases    = []string{"request_entity_too_large", "request entity too large"}
	requestEntityTooLargeValidators = []func(string) bool{
		func(s string) bool { return streams.From(requestEntityTooLargeAliases).Contains(strings.ToLower(s)) },
	}
)

// NewRequestEntityTooLarge creates a new `Error` and sets the error code to 413 Request Entity Too Large (ErrorType_REQUEST_ENTITY_TOO_LARGE) and title to the default for that status.
func NewRequestEntityTooLarge(msg string, errs ...error) *Error {
	return newError(1, http.StatusRequestEntityTooLarge, http.StatusText(http.StatusRequestEntityTooLarge), msg, errs...)
}

// WrapRequestEntityTooLarge wraps an error and sets the error code to 413 Request Entity Too Large (ErrorType_REQUEST_ENTITY_TOO_LARGE) and title to the default for that status.
func WrapRequestEntityTooLarge(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusRequestEntityTooLarge, msg...)
}

// WrapRequestEntityTooLargef wraps an error and sets the error code to 413 Request Entity Too Large (ErrorType_REQUEST_ENTITY_TOO_LARGE) using as message the string format result
// of `format` and `args
func WrapRequestEntityTooLargef(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusRequestEntityTooLarge, fmt.Sprintf(format, args...))
}

// IsRequestEntityTooLarge indicates whether the request payload was too large
func IsRequestEntityTooLarge(err error) bool {
	return IsErrorType(http.StatusRequestEntityTooLarge, err)
}

func isRequestEntityTooLarge(msg string) bool {
	return streams.From(requestEntityTooLargeValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    500 Unhandled
// ******************************************************

var (
	unhandledAliases    = []string{"unhandled", "unhandled", "internal server error"}
	unhandledValidators = []func(string) bool{
		func(s string) bool { return streams.From(unhandledAliases).Contains(strings.ToLower(s)) },
	}
)

// NewUnhandled creates a new `Error` and sets the error code to 500 Unhandled (ErrorType_UNHANDLED) and title to the default for that status.
func NewUnhandled(msg string, errs ...error) *Error {
	return newError(1, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), msg, errs...)
}

// WrapUnhandled wraps an error and sets the error code to 500 Unhandled (ErrorType_UNHANDLED) and title to the default for that status.
func WrapUnhandled(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusInternalServerError, msg...)
}

// WrapUnhandledf wraps an error and sets the error code to 500 Unhandled (ErrorType_UNHANDLED) using as message the string format result
// of `format` and `args
func WrapUnhandledf(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusInternalServerError, fmt.Sprintf(format, args...))
}

// IsUnhandled indicates whether the error points to unhandled server-side error
func IsUnhandled(err error) bool {
	return IsErrorType(http.StatusRequestEntityTooLarge, err)
}

func isUnhandled(msg string) bool {
	return streams.From(unhandledValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    501 Not Implemented
// ******************************************************

var (
	notImplementedAliases    = []string{"not_implemented", "not implemented"}
	notImplementedValidators = []func(string) bool{
		func(s string) bool { return streams.From(notImplementedAliases).Contains(strings.ToLower(s)) },
	}
)

// NewNotImplemented creates a new `Error` and sets the error code to 501 Not Implemented (ErrorType_NOT_IMPLEMENTED) and title to the default for that status.
func NewNotImplemented(msg string, errs ...error) *Error {
	return newError(1, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), msg, errs...)
}

// WrapNotImplemented wraps an error and sets the error code to 501 Not Implemented (ErrorType_NOT_IMPLEMENTED) and title to the default for that status.
func WrapNotImplemented(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusNotImplemented, msg...)
}

// WrapNotImplementedf wraps an error and sets the error code to 501 Not Implemented (ErrorType_NOT_IMPLEMENTED) using as message the string format result
// of `format` and `args
func WrapNotImplementedf(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusNotImplemented, fmt.Sprintf(format, args...))
}

// IsNotImplemented indicates whether the error points to method that has not been implemented yet
func IsNotImplemented(err error) bool {
	return IsErrorType(http.StatusNotImplemented, err)
}

func isNotImplemented(msg string) bool {
	return streams.From(notImplementedValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}

// ******************************************************
//    502 Bad Gateway
// ******************************************************

var (
	badGatewayAliases    = []string{"bad_gateway", "bad gateway"}
	badGatewayValidators = []func(string) bool{
		func(s string) bool { return streams.From(badGatewayAliases).Contains(strings.ToLower(s)) },
	}
)

// NewBadGateway creates a new `Error` and sets the error code to 502 Bad Gateway (ErrorType_BAD_GATEWAY) and title to the default for that status.
func NewBadGateway(msg string, errs ...error) *Error {
	return newError(1, http.StatusBadGateway, http.StatusText(http.StatusBadGateway), msg, errs...)
}

// WrapBadGateway wraps an error and sets the error code to 502 Bad Gateway (ErrorType_BAD_GATEWAY) and title to the default for that status.
func WrapBadGateway(err error, msg ...string) *Error {
	return wrapWithCode(1, err, http.StatusBadGateway, msg...)
}

// WrapBadGatewayf wraps an error and sets the error code to 502 Bad Gateway (ErrorType_BAD_GATEWAY) using as message the string format result
// of `format` and `args
func WrapBadGatewayf(err error, format string, args ...interface{}) *Error {
	return wrapWithCode(1, err, http.StatusBadGateway, fmt.Sprintf(format, args...))
}

// IsBadGateway indicates whether the error points to a bad gateway which occurred from a service-to-service communication
func IsBadGateway(err error) bool {
	return IsErrorType(http.StatusBadGateway, err)
}

func isBadGateway(msg string) bool {
	return streams.From(badGatewayValidators).
		AnyMatch(func(i interface{}) bool {
			s := i.(func(string) bool)
			return s(msg)
		})
}
