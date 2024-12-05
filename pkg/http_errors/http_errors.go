package http_errors

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	ErrBadRequest       = "Bad request"
	ErrAlreadyExists    = "Already exists"
	ErrNoSuchUser       = "User not found"
	ErrWrongCredentials = "Wrong Credentials"
	ErrNotFound         = "Not Found"
	ErrUnauthorized     = "Unauthorized"
	ErrForbidden        = "Forbidden"
	ErrBadQueryParams   = "Invalid query params"
	ErrRequestTimeout   = "Request Timeout"
	ErrInvalidEmail     = "Invalid email"
	ErrInvalidPassword  = "Invalid password"
	ErrInvalidField     = "Invalid field"
)

var (
	BadRequest            = errors.New("Bad request")
	WrongCredentials      = errors.New("Wrong Credentials")
	NotFound              = errors.New("Not Found")
	Unauthorized          = errors.New("Unauthorized")
	Forbidden             = errors.New("Forbidden")
	UserNotFound          = errors.New("User not found")
	PermissionDenied      = errors.New("Permission Denied")
	ExpiredCSRFError      = errors.New("Expired CSRF token")
	WrongCSRFToken        = errors.New("Wrong CSRF token")
	CSRFNotPresented      = errors.New("CSRF not presented")
	NotRequiredFields     = errors.New("No such required fields")
	BadQueryParams        = errors.New("Invalid query params")
	InternalServerError   = errors.New("Internal Server Error")
	RequestTimeoutError   = errors.New("Request Timeout")
	ExistsEmailError      = errors.New("User with given email already exists")
	InvalidJWTToken       = errors.New("Invalid JWT token")
	InvalidJWTClaims      = errors.New("Invalid JWT claims")
	NotAllowedImageHeader = errors.New("Not allowed image header")
	NoCookie              = errors.New("not found cookie header")
)

type HTTPErr interface {
	StatusCode() int
	Error() string
	Causes() interface{}
	ErrBody() HTTPError
}

type HTTPError struct {
	ErrStatusCode int         `json:"status,omitempty"`
	ErrError      string      `json:"error,omitempty"`
	ErrCauses     interface{} `json:"causes,omitempty"`
}

func (e HTTPError) ErrBody() HTTPError {
	return e
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("status: %d, error: %s, causes: %v", e.ErrStatusCode, e.ErrError, e.ErrCauses)
}

func (e HTTPError) StatusCode() int {
	return e.ErrStatusCode
}

func (e HTTPError) Causes() interface{} {
	return e.ErrCauses
}

func NewHTTPError(status int, err string, causes interface{}) HTTPErr {
	return HTTPError{
		ErrStatusCode: status,
		ErrError:      err,
		ErrCauses:     causes,
	}
}

func NewInternalServerError(causes interface{}) HTTPErr {
	result := HTTPError{
		ErrStatusCode: http.StatusInternalServerError,
		ErrError:      InternalServerError.Error(),
		ErrCauses:     causes,
	}

	return result
}

func ParseErrors(err error) HTTPErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewHTTPError(http.StatusNotFound, ErrNotFound, nil)

	case errors.Is(err, context.DeadlineExceeded):
		return NewHTTPError(http.StatusRequestTimeout, ErrRequestTimeout, nil)

	case errors.Is(err, Unauthorized):
		return NewHTTPError(http.StatusUnauthorized, ErrUnauthorized, nil)

	case errors.Is(err, WrongCredentials):
		return NewHTTPError(http.StatusUnauthorized, ErrUnauthorized, nil)

	case errors.Is(err, UserNotFound):
		return NewHTTPError(http.StatusBadRequest, ErrNotFound, "Invalid login data")

	case strings.Contains(strings.ToLower(err.Error()), "sqlstate"):
		return parseSqlErrors(err)

	case strings.Contains(strings.ToLower(err.Error()), "field validation"):
		return parseValidatorError(err)

	case strings.Contains(strings.ToLower(err.Error()), "unmarshal"):
		return NewHTTPError(http.StatusBadRequest, ErrBadRequest, err)

	case strings.Contains(strings.ToLower(err.Error()), "uuid"):
		return NewHTTPError(http.StatusBadRequest, ErrBadRequest, err)

	case strings.Contains(strings.ToLower(err.Error()), "cookie"):
		return NewHTTPError(http.StatusUnauthorized, ErrUnauthorized, err)

	case strings.Contains(strings.ToLower(err.Error()), "token"):
		return NewHTTPError(http.StatusUnauthorized, ErrUnauthorized, err)

	case strings.Contains(strings.ToLower(err.Error()), "bcrypt"):
		return NewHTTPError(http.StatusBadRequest, ErrBadRequest, nil)

	case strings.Contains(strings.ToLower(err.Error()), "email_1 dup key"):
		return NewHTTPError(http.StatusBadRequest, ErrBadRequest, err)

	default:
		if restErr, ok := err.(HTTPErr); ok {
			return restErr
		}
		return NewInternalServerError(err)
	}
}

func parseSqlErrors(err error) HTTPErr {
	return NewHTTPError(http.StatusBadRequest, ErrBadRequest, err)
}

func parseValidatorError(err error) HTTPErr {
	if strings.Contains(err.Error(), "Password") {
		return NewHTTPError(http.StatusBadRequest, ErrInvalidPassword, err)
	}

	if strings.Contains(err.Error(), "Email") {
		return NewHTTPError(http.StatusBadRequest, ErrInvalidEmail, err)
	}

	return NewHTTPError(http.StatusBadRequest, ErrInvalidField, err)
}

func ErrorResponse(err error) (int, interface{}) {
	return ParseErrors(err).StatusCode(), ParseErrors(err)
}

func NewErrorResponse(c *gin.Context, err error) {
	restErr := ParseErrors(err)
	c.JSON(restErr.StatusCode(), restErr.ErrBody())
}
