package api

import "net/http"

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

// Error implements the Error interface
func (e Error) Error() string {
	return e.Err
}

func ErrUnAuthorized() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Err:  "unauthorized request",
	}
}

func ErrBadRequest() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid JSON request",
	}
}

func ErrResourceNotFound(res string) Error {
	return Error{
		Code: http.StatusNotFound,
		Err:  res + "resource not found",
	}
}

func ErrInvalidID() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid id given",
	}
}