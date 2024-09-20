package product

import (
	"acme/model"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	Close()
}
