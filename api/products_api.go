package api

import (
	"acme/service"
	"encoding/json"
	"net/http"
)

type ProductAPI struct {
	productService *service.ProductService
}

func NewProductPI(productService *service.ProductService) *ProductAPI {
	return &ProductAPI{
		productService: productService,
	}
}

func (api *ProductAPI) GetProducts(writer http.ResponseWriter, request *http.Request) {

	products, err := api.productService.GetProducts()

	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(products)

}
