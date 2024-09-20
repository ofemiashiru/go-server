package product

import (
	"acme/model"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	AddProduct(product []model.Product) (id int, err error)
	Close()
}
