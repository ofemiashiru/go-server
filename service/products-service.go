package service

import (
	"acme/model"
	"acme/repository/product"
	"errors"
	"fmt"
)

type ProductService struct {
	repository product.ProductRepository
}

// NewProductService creates a new instance of ProductService.
func NewProductService(repo product.ProductRepository) *ProductService {
	return &ProductService{
		repository: repo,
	}
}

func (s *ProductService) GetProducts() ([]model.Product, error) {
	products, err := s.repository.GetProducts()

	if err != nil {
		fmt.Println("Error getting products from DB:", err)
		return nil, errors.New("there was an error getting the users from the database")
	}

	return products, nil

}
