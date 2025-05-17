package pkg

import (
	"errors"
	"net/http"
)

type ApiError interface {
	GetStatus() int
	GetResponse() ErrorResponse
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ToApiError(err Error) ApiError {
	var notFoundError isNotFound
	if errors.As(err, &notFoundError) {
		return NewNotFoundApiError(notFoundError)
	}

	var forbiddenError isForbidden
	if errors.As(err, &forbiddenError) {
		return NewForbiddenApiError(forbiddenError)
	}

	var unauthorizedError isUnauthorized
	if errors.As(err, &unauthorizedError) {
		return NewUnauthorizedApiError(unauthorizedError)
	}

	var fatalError isFatal
	if errors.As(err, &fatalError) {
		return NewInternalServerApiError(fatalError)
	}

	return NewBadRequestApiError(err)
}

// NotFoundApiError ----------------------------------------------------------------------------------------------------
type NotFoundApiError struct {
	status  int
	message string
}

func NewNotFoundApiError(err Error) NotFoundApiError {
	return NotFoundApiError{
		status:  http.StatusNotFound,
		message: err.getMessage(),
	}
}

func (ae NotFoundApiError) GetStatus() int {
	return ae.status
}

func (ae NotFoundApiError) GetResponse() ErrorResponse {
	return ErrorResponse{
		Code:    ae.status,
		Message: ae.message,
	}
}

// BadRequestApiError --------------------------------------------------------------------------------------------------
type BadRequestApiError struct {
	status  int
	message string
}

func NewBadRequestApiError(err Error) BadRequestApiError {
	return BadRequestApiError{
		status:  http.StatusBadRequest,
		message: err.getMessage(),
	}
}

func (ae BadRequestApiError) GetStatus() int {
	return ae.status
}

func (ae BadRequestApiError) GetResponse() ErrorResponse {
	return ErrorResponse{
		Code:    ae.status,
		Message: ae.message,
	}
}

// ForbiddenApiError ---------------------------------------------------------------------------------------------------
type ForbiddenApiError struct {
	status  int
	message string
}

func NewForbiddenApiError(err Error) ForbiddenApiError {
	return ForbiddenApiError{
		status:  http.StatusForbidden,
		message: err.getMessage(),
	}
}

func (ae ForbiddenApiError) GetStatus() int {
	return ae.status
}

func (ae ForbiddenApiError) GetResponse() ErrorResponse {
	return ErrorResponse{
		Code:    ae.status,
		Message: ae.message,
	}
}

// UnauthorizedApiError ---------------------------------------------------------------------------------------------------
type UnauthorizedApiError struct {
	status  int
	message string
}

func NewUnauthorizedApiError(err Error) UnauthorizedApiError {
	return UnauthorizedApiError{
		status:  http.StatusUnauthorized,
		message: err.getMessage(),
	}
}

func (ae UnauthorizedApiError) GetStatus() int {
	return ae.status
}

func (ae UnauthorizedApiError) GetResponse() ErrorResponse {
	return ErrorResponse{
		Code:    ae.status,
		Message: ae.message,
	}
}

// InternalServerApiError ----------------------------------------------------------------------------------------------
type InternalServerApiError struct {
	status  int
	message string
}

func NewInternalServerApiError(err Error) InternalServerApiError {
	return InternalServerApiError{
		status:  http.StatusInternalServerError,
		message: err.getMessage(),
	}
}

func (ae InternalServerApiError) GetStatus() int {
	return ae.status
}

func (ae InternalServerApiError) GetResponse() ErrorResponse {
	return ErrorResponse{
		Code:    ae.status,
		Message: ae.message,
	}
}
