package api

import (
	"acme/model"
	"acme/service"
	"encoding/json"
	"fmt"
	"io"
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

func (api *ProductAPI) CreateProudct(writer http.ResponseWriter, request *http.Request) {

	product, err := decodeProduct(request.Body)

	if err != nil {
		http.Error(writer, "Bad Request Body", http.StatusBadRequest)
		return
	}

	newProductSlice := []model.Product{product}

	id, err := api.productService.CreateProduct(newProductSlice)

	if err != nil {
		http.Error(writer, "Product not created", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "Product created successfully: %d", id)

}

func decodeProduct(body io.ReadCloser) (product model.Product, err error) {

	err = json.NewDecoder(body).Decode(&product)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return model.Product{}, err
	}

	return product, nil
}
