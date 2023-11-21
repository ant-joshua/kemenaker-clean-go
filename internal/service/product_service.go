package service

import (
	"clean_go/internal/domains"
	"clean_go/internal/kafka"
	"fmt"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	// db connection
	db    *gorm.DB
	kafka kafka.CustomKafka
}

func (p ProductServiceImpl) GetAllProduct(filter domains.GetAllProductFilter) ([]domains.Product, error) {
	var productList []domains.Product

	productList = make([]domains.Product, 0)

	productList = append(productList, domains.Product{
		ID:          1,
		Name:        "Kopi",
		Price:       1000,
		Description: "ini kopi",
		//Categories:  nil,
	})

	productList = append(productList, domains.Product{
		ID:          2,
		Name:        "Teh",
		Price:       1000,
		Description: "ini Teh",
		//Categories:  nil,
	})

	return productList, nil

}

func (p ProductServiceImpl) GetProductById(id int) (*domains.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServiceImpl) CreateProduct(request domains.CreateProductRequest) (*domains.Product, error) {

	var newProduct domains.Product

	newProduct.Name = request.Name
	newProduct.Price = request.Price
	newProduct.Description = request.Description

	err := p.db.Create(&newProduct).Error

	if err != nil {

		fmt.Printf("error db %v", err)

		return nil, err
	}

	p.kafka.PublishMessage()

	return &newProduct, nil
}

func (p ProductServiceImpl) UpdateProduct(request domains.UpdateProductRequest) (*domains.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServiceImpl) DeleteProducts(ids []int) error {
	//TODO implement me
	panic("implement me")
}

type ProductService interface {
	GetAllProduct(filter domains.GetAllProductFilter) ([]domains.Product, error)
	GetProductById(id int) (*domains.Product, error)
	CreateProduct(request domains.CreateProductRequest) (*domains.Product, error)
	UpdateProduct(request domains.UpdateProductRequest) (*domains.Product, error)
	DeleteProducts(ids []int) error
}

func NewProductService(db *gorm.DB) ProductService {
	return &ProductServiceImpl{
		db: db,
	}
}
