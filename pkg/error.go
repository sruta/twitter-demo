package pkg

import (
	"fmt"
)

func formatMessage(message string, original error) string {
	/* Uncomment this to debug the original cause of the error
	if original != nil {
		return message + fmt.Sprintf(", original: %s", original.Error())
	}
	*/
	return message
}

type Error interface {
	getMessage() string
	Error() string
}

// NotFoundError -------------------------------------------------------------------------------------------------------
type NotFoundError struct {
	Message string
	error
}

func NewNotFoundError(message string, original error) NotFoundError {
	return NotFoundError{
		Message: message,
		error:   original,
	}
}

func NewDBNotFoundError(message string, original error) NotFoundError {
	message = fmt.Sprintf("%s doesn't exist", message)
	return NotFoundError{
		Message: message,
		error:   original,
	}
}

func IsNotFound(err error) bool {
	typedError, ok := err.(isNotFound)
	return ok && typedError.isNotFoundError()
}

type isNotFound interface {
	isNotFoundError() bool
	getMessage() string
	Error() string
}

func (e NotFoundError) isNotFoundError() bool {
	return true
}

func (e NotFoundError) getMessage() string {
	return formatMessage(e.Message, e.error)
}

// ForbiddenError ------------------------------------------------------------------------------------------------------
type ForbiddenError struct {
	Message string
	error
}

func NewForbiddenError(message string, original error) ForbiddenError {
	return ForbiddenError{
		Message: message,
		error:   original,
	}
}

func IsForbidden(err error) bool {
	typedError, ok := err.(isForbidden)
	return ok && typedError.isForbiddenError()
}

type isForbidden interface {
	isForbiddenError() bool
	getMessage() string
	Error() string
}

func (e ForbiddenError) isForbiddenError() bool {
	return true
}

func (e ForbiddenError) getMessage() string {
	return formatMessage(e.Message, e.error)
}

// UnauthorizedError ---------------------------------------------------------------------------------------------------
type UnauthorizedError struct {
	Message string
	error
}

func NewUnauthorizedError(message string, original error) UnauthorizedError {
	return UnauthorizedError{
		Message: message,
		error:   original,
	}
}

func NewNotLoggedInUnauthorizedError(original error) UnauthorizedError {
	return NewUnauthorizedError("user not logged in", original)
}

func IsUnauthorized(err error) bool {
	typedError, ok := err.(isUnauthorized)
	return ok && typedError.isUnauthorizedError()
}

type isUnauthorized interface {
	isUnauthorizedError() bool
	getMessage() string
	Error() string
}

func (e UnauthorizedError) isUnauthorizedError() bool {
	return true
}

func (e UnauthorizedError) getMessage() string {
	return formatMessage(e.Message, e.error)
}

// GenericError --------------------------------------------------------------------------------------------------------
type GenericError struct {
	Message string
	error
}

func NewGenericError(message string, original error) GenericError {
	return GenericError{
		Message: message,
		error:   original,
	}
}

func NewInvalidBodyGenericError(original error) GenericError {
	return NewGenericError("invalid request body", original)
}

func NewInvalidIDGenericError(original error) GenericError {
	return NewGenericError("invalid request id", original)
}

func IsGeneric(err error) bool {
	typedError, ok := err.(isGeneric)
	return ok && typedError.isGenericError()
}

type isGeneric interface {
	isGenericError() bool
	getMessage() string
	Error() string
}

func (e GenericError) isGenericError() bool {
	return true
}

func (e GenericError) getMessage() string {
	return formatMessage(e.Message, e.error)
}

// FatalError ----------------------------------------------------------------------------------------------------------
type FatalError struct {
	Message string
	error
}

func NewFatalError(message string, original error) FatalError {
	return FatalError{
		Message: message,
		error:   original,
	}
}

func NewDBFatalError(message string, original error) FatalError {
	return NewFatalError(fmt.Sprintf("could not %s db", message), original)
}

func NewDBScanFatalError(message string, original error) FatalError {
	return NewFatalError(fmt.Sprintf("could not deserialize %s from db", message), original)
}

func IsFatal(err error) bool {
	typedError, ok := err.(isFatal)
	return ok && typedError.isFatalError()
}

type isFatal interface {
	isFatalError() bool
	getMessage() string
	Error() string
}

func (e FatalError) isFatalError() bool {
	return true
}

func (e FatalError) getMessage() string {
	return formatMessage(e.Message, e.error)
}
