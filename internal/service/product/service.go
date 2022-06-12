package product

import (
	"errors"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(id int) (*Product, error) {
	if id >= len(allProducts) || id < 0 {
		return nil, errors.New("Invalid arguments")
	}

	return &allProducts[id], nil
}
