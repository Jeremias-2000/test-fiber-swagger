package routes

import (
	"github.com/test-fiber/models"
)

type Product struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID,Name: productModel.Name,SerialNumber: productModel.SerialNumber}
}



