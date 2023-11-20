package commons

import "github.com/labstack/echo/v4"

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type SuccessResponseProps struct {
	Success *bool
	Message *string
	Code    *int
	Data    interface{}
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
}

type CustomResponse struct {
}

func (c CustomResponse) Success(ctx echo.Context, props SuccessResponseProps) error {
	var message string = "success"
	var code int = 200
	var data interface{} = nil

	if props.Code != nil {
		code = *props.Code
	}

	if props.Data != nil {
		data = props.Data
	}

	return ctx.JSON(code, SuccessResponse{
		Success: true,
		Message: message,
		Code:    code,
		Data:    data,
	})
}

func (c CustomResponse) BadRequest(ctx echo.Context, errors map[string]interface{}) error {

	//var errorMap = make(map[string]interface{})

	//if errors != nil {
	//	for _, err := range errors {
	//		for key, val := range err {
	//			errorMap[key] = val
	//		}
	//	}
	//}

	return ctx.JSON(400, ErrorResponse{
		Success: false,
		Message: "Bad Request",
		Code:    400,
		Errors:  errors,
	})
}
