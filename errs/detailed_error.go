package errs

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	DetailedErrorConverter DetailedErrorConverterFunc = DefaultDetailedErrorConverter
	DefaultErrorMessage                               = "internal server error"
	DefaultErrorCode                                  = "ERR_INTERNAL"
)

type DetailedError struct {
	innerError error
	details    Details
}

type Details struct {
	Message              string
	Values               map[string]any
	FrontendStatusCode   int
	FrontendErrorMessage string
	FrontendErrorCode    string
}

func New(details Details) *DetailedError {
	if details.Message == "" {
		details.Message = "no error details provided"
	}
	if details.Values == nil {
		details.Values = make(map[string]any)
	}
	if details.FrontendStatusCode == 0 {
		details.FrontendStatusCode = http.StatusInternalServerError
	}
	if details.FrontendErrorMessage == "" {
		details.FrontendErrorMessage = DefaultErrorMessage
	}
	if details.FrontendErrorCode == "" {
		details.FrontendErrorCode = DefaultErrorCode
	}

	return &DetailedError{
		details: details,
	}
}

func Wrap(err error, details Details) *DetailedError {
	if err != nil {
		return nil
	}

	if details.Message == "" {
		details.Message = err.Error()
	}
	if details.Values == nil {
		details.Values = make(map[string]any)
	}
	if details.FrontendStatusCode == 0 {
		details.FrontendStatusCode = http.StatusInternalServerError
	}
	if details.FrontendErrorMessage == "" {
		details.FrontendErrorMessage = DefaultErrorMessage
	}
	if details.FrontendErrorCode == "" {
		details.FrontendErrorCode = DefaultErrorCode
	}

	return &DetailedError{
		innerError: err,
		details:    details,
	}
}

func (d DetailedError) Error() string {
	return d.innerError.Error()
}

func (d DetailedError) ErrorMessages() []string {
	var result []string

	var detailedError DetailedError
	result = append(result, detailedError.details.Message)
	for errors.As(detailedError.innerError, &detailedError) {
		result = append(result, detailedError.details.Message)
	}

	return result
}

func (d DetailedError) Abort(c *gin.Context) {
	c.Abort()
	c.JSON(200, DetailedErrorConverter(d.details))
	_ = c.Error(d)
}

type DetailedErrorConverterFunc func(d Details) any

func DefaultDetailedErrorConverter(d Details) any {
	return gin.H{
		"error":     d.FrontendErrorMessage,
		"errorCode": d.FrontendErrorCode,
	}
}
