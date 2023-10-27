package utils

import "net/http"

type APIError struct {
	Code    int
	Message string
}

var (
	ErrDataNotFound = APIError{
		Code:    http.StatusNotFound,
		Message: "Data Not Found",
	}

	ErrInternalServerError = APIError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}

	ErrBadRequest = APIError{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
	}

	SuccessAddData = APIError{
		Code:    http.StatusOK,
		Message: "Data Added Successfully",
	}
)
