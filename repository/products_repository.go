package repository

import (
	"TokoBelanja/model/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type productsRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productsRepository {
	return &productsRepository{db}
}

func (r *productsRepository) Save(product *entity.Product) (*entity.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productsRepository) GetAll() ([]entity.Product, error) {
	var product []entity.Product
	err := r.db.Find(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productsRepository) FindById(id int) (*entity.Product, error) {
	var product *entity.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product with id %d not found", id)
		}
		return nil, err
	}
	return product, nil
}

func (r *productsRepository) Update(id int, product *entity.Product) (*entity.Product, error) {
	err := r.db.Where("id = ?", id).Updates(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productsRepository) Delete(id int) error {
	var product *entity.Product
	err := r.db.Where("id = ?", id).Delete(&product).Error
	return err
}
