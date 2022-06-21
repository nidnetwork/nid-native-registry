package helpers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ErrorLog -
func ErrorLog(status int, c *gin.Context, err error, code string) {
	log.Printf("Error %d: %s %s, %s - %s", status, c.Request.Method, c.Request.URL.Path, code, err.Error())
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// ValidationDetail -
type ValidationDetail struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

// ValidationError -
type ValidationError struct {
	Code    string             `json:"code" example:"ValidationError"`
	Message string             `json:"message" example:"invalid request"`
	Errors  []ValidationDetail `json:"errors"`
}

// BadRequestError -
type BadRequestError struct {
	Code    string `json:"code" example:"BadRequestError"`
	Message string `json:"message" example:"status bad request"`
}

// ServerError -
type ServerError struct {
	Code    string `json:"code" example:"ServerError"`
	Message string `json:"message" example:"internal error"`
}

// Descriptive -
func Descriptive(verr validator.ValidationErrors) []ValidationDetail {
	errs := []ValidationDetail{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationDetail{Field: f.Field(), Reason: err})
	}

	return errs
}

// NewValidationError -
func NewValidationError(ctx *gin.Context, err error) {
	status := http.StatusBadRequest
	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		er := ValidationError{
			Code:    "ValidationError",
			Message: err.Error(),
			Errors:  Descriptive(verr),
		}
		ctx.JSON(status, er)
	} else {
		ctx.JSON(status, BadRequestError{
			Code:    "ValidationError",
			Message: err.Error(),
		})
	}
}

// NewBadRequestError -
func NewBadRequestError(ctx *gin.Context, err error, code string) {
	status := http.StatusBadRequest
	ErrorLog(status, ctx, err, code)
	ctx.JSON(status, BadRequestError{
		Code:    code,
		Message: err.Error(),
	})
}

// NewServerError -
func NewServerError(ctx *gin.Context, err error, code string) {
	status := http.StatusInternalServerError
	ErrorLog(status, ctx, err, code)
	ctx.JSON(status, ServerError{
		Code:    code,
		Message: err.Error(),
	})
}

// NewUnauthorizedError -
func NewUnauthorizedError(ctx *gin.Context, err error, code string) {
	status := http.StatusUnauthorized
	ErrorLog(status, ctx, err, code)
	ctx.JSON(status, BadRequestError{
		Code:    code,
		Message: err.Error(),
	})
}
