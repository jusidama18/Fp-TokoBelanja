package service

import (
	"TokoBelanja/model/entity"
	"TokoBelanja/model/input"
	"TokoBelanja/repository"
	"errors"
	"fmt"
)

type ProductService interface {
	Create(role string, input *input.ProductCreateInput) (*entity.Product, error)
	GetAll(role string) ([]entity.Product, error)
	Put(role string, id int, input *input.ProductPutInput) (*entity.Product, error)
	Delete(role string, id int) error
}

type productService struct {
	repo repository.ProductsRepository
}

func NewProductService(repo repository.ProductsRepository) *productService {
	return &productService{repo: repo}
}

func (s *productService) Create(role string, input *input.ProductCreateInput) (*entity.Product, error) {
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
	return s.repo.Create(product)
}

func (s *productService) GetAll(role string) ([]entity.Product, error) {
	if role != "admin" {
		return nil, errors.New("you are not admin")
	}

	return s.repo.GetAll()
}

func (s *productService) Put(role string, id int, input *input.ProductPutInput) (*entity.Product, error) {
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

	check, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	if check.ID == 0 {
		return nil, fmt.Errorf("product with id %d not found", id)
	}

	new_product := &entity.Product{
		Title:      input.Title,
		Price:      input.Price,
		Stock:      input.Stock,
		CategoryID: input.CategoryID,
	}
	return s.repo.Update(id, new_product)
}

func (s *productService) Delete(role string, id int) error {
	if role != "admin" {
		return errors.New("you are not admin")
	}

	check, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	if check.ID == 0 {
		return fmt.Errorf("product with id %d not found", id)
	}

	return s.repo.Delete(id)
}
