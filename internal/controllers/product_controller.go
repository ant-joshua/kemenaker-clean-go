package controllers

import (
	"clean_go/internal/commons"
	"clean_go/internal/domains"
	"clean_go/internal/service"
	"github.com/labstack/echo/v4"
)

type ProductHttpControllerImpl struct {
	service service.ProductService
}

func (p ProductHttpControllerImpl) Routes(e *echo.Echo) {
	//TODO implement me
	e.GET("/products", p.GetAllProduct)
	e.POST("/products", p.CreateProduct)
	e.GET("/products/{id}", p.GetProductById)
	e.PUT("/products/:id", p.UpdateProduct)
	e.DELETE("/products", p.DeleteProducts)
}

func (p ProductHttpControllerImpl) GetAllProduct(ctx echo.Context) error {

	//validate := validator.New(validator.WithRequiredStructEnabled())
	filterRequest := new(domains.GetAllProductFilter)
	response := new(commons.CustomResponse)

	if err := ctx.Bind(filterRequest); err != nil {
		//var error map[string]interface{}

		return response.BadRequest(ctx, map[string]interface{}{
			"page":  "is required",
			"limit": "is required",
		})
	}

	//if err := validate.Struct(filterRequest); err != nil {
	//
	//	fmt.Printf("%v", err)
	//	return err
	//}

	products, err := p.service.GetAllProduct(*filterRequest)

	if err != nil {
		return ctx.JSON(500, echo.Map{
			"message": "internal server error",
			"errors":  nil,
		})
	}

	props := new(commons.SuccessResponseProps)

	props.Data = products

	return response.Success(ctx, *props)

}

func (p ProductHttpControllerImpl) GetProductById(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductHttpControllerImpl) CreateProduct(ctx echo.Context) error {
	createRequest := new(domains.CreateProductRequest)

	response := new(commons.CustomResponse)

	if err := ctx.Bind(createRequest); err != nil {
		return response.BadRequest(ctx, map[string]interface{}{})
	}

	product, err := p.service.CreateProduct(*createRequest)

	if err != nil {
		return response.BadRequest(ctx, map[string]interface{}{})
	}

	props := new(commons.SuccessResponseProps)
	props.Data = product

	return response.Success(ctx, *props)

}

func (p ProductHttpControllerImpl) UpdateProduct(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductHttpControllerImpl) DeleteProducts(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

type ProductHttpController interface {
	GetAllProduct(ctx echo.Context) error
	GetProductById(ctx echo.Context) error
	CreateProduct(ctx echo.Context) error
	UpdateProduct(ctx echo.Context) error
	DeleteProducts(ctx echo.Context) error
	Routes(echo *echo.Echo)
}

func NewProductHttpController(service service.ProductService) ProductHttpController {
	return ProductHttpControllerImpl{
		service: service,
	}
}
