package memory

import (
	"C_12S-practica_arquitectura/internal/products"
	"encoding/json"
	"errors"
	"os"
)

type Service struct {
	products []products.Product
}

func NewService(path string) (*Service, error) {
	byteFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var products []products.Product
	err = json.Unmarshal(byteFile, &products)
	if err != nil {
		return nil, err
	}
	return &Service{products: products}, nil
}

func (s *Service) GetByID(id int) (products.Product, error) {
	for _, product := range s.products {
		if product.ID == id {
			return product, nil
		}
	}
	return products.Product{}, errors.New("not found")
}

func (s *Service) Modify(id int, product products.Product) (products.Product, error) {
	for i, value := range s.products {
		if value.ID == id {
			s.products = append(s.products[:i], s.products[i+1:]...)
			s.products = append(s.products, product)
			return product, nil
		}
	}
	return products.Product{}, errors.New("not found")
}
