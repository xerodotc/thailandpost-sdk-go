package thailandpost

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrBadRequest      = errors.New("bad request")
	ErrServerError     = errors.New("server error")
	ErrForbidden       = errors.New("forbidden")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrTooManyRequests = errors.New("too many requests")

	ErrUnsuccessful = errors.New("unsuccessful")

	ErrTooMuchTrackings      = errors.New("too much tracking numbers, limited to under 100 tracking numbers")
	ErrInvalidTrackingNumber = errors.New("invalid tracking number")

	ErrInvalidBearerToken = errors.New("invalid bearer token")

	ErrNoLastRequest = errors.New("no last succesful request")
)

type InvalidTrackingNumbersError struct {
	InvalidTrackingNumbers []string
}

func (e InvalidTrackingNumbersError) Error() string {
	return fmt.Sprintf("invalid tracking numbers: %s", strings.Join(e.InvalidTrackingNumbers, ", "))
}

func (e InvalidTrackingNumbersError) Unwrap() error {
	return ErrInvalidTrackingNumber
}

type UnsuccessfulError struct {
	Message string
}

func (e UnsuccessfulError) Error() string {
	return e.Message
}

func (e UnsuccessfulError) Unwrap() error {
	return ErrUnsuccessful
}

func httpStatusCodeToError(code int) error {
	switch code {
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusInternalServerError:
		return ErrServerError
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusTooManyRequests:
		return ErrTooManyRequests
	}

	return nil
}
