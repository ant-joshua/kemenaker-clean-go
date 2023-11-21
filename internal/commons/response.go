package commons

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response[T any] struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    T           `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func (r *Response[T]) SetSuccess(success bool) *Response[T] {
	r.Success = success

	return r
}

func (r *Response[T]) SetMessage(message string) *Response[T] {
	r.Message = message

	return r
}

func (r *Response[T]) SetCode(code int) *Response[T] {
	r.Code = code

	return r
}

func (r *Response[T]) SetData(data T) *Response[T] {
	r.Data = data

	return r
}

func (r *Response[T]) SetErrors(errors interface{}) *Response[T] {
	r.Errors = errors

	return r
}

func NewSuccessResponse[T any](ctx echo.Context, message string, data *T, code *int) error {
	if message == "" {
		message = "Success"
	}
	if code == nil {
		code = new(int)
		*code = http.StatusOK
	}

	return ctx.JSON(*code, Response[T]{
		Success: true,
		Message: message,
		Code:    *code,
		Data:    *data,
	})
}

func NewErrorResponse(ctx echo.Context, message string, code int, errors interface{}) error {
	if message == "" {
		message = "Error"
	}
	if code == 0 {
		code = http.StatusBadRequest
	}

	return ctx.JSON(code, Response[interface{}]{
		Success: false,
		Message: message,
		Code:    code,
		Errors:  errors,
	})
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewBadRequestResponse(ctx echo.Context, err error) error {
	var errors []map[string]interface{}
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, map[string]interface{}{
			"field":   e.Field(),
			"message": e.Tag(),
		})
	}
	return NewErrorResponse(ctx, "Bad Request", http.StatusBadRequest, errors)
}
