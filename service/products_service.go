package service

import (
	"TokoBelanja/model/entity"
	"TokoBelanja/model/input"
	"TokoBelanja/repository"
	"errors"
)

type productService struct {
	repo repository.ProductsRepository
}

func NewProductService(repo repository.ProductsRepository) *productService {
	return &productService{repo: repo}
}

func (s *productService) Create(role string, input input.ProductCreateInput) (*entity.Product, error) {
	if role != "admin" {
		return nil, errors.New("you are not admin")
	}

	if input.Title == "" {
		return nil, errors.New("title should not empty")
	}

	if input.Stock < 5 {
		return nil, errors.New("stock should not less than 5")
	}

	if input.Price > 50000000 || input.Price < 0 {
		return nil, errors.New("price should less than 50.000.000 or more than 0")
	}

	product := &entity.Product{
		Title:      input.Title,
		Price:      input.Price,
		Stock:      input.Stock,
		CategoryID: input.CategoryID,
	}
	return s.repo.Save(product)
}


