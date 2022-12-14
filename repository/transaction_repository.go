package repository

import (
	"TokoBelanja/model/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(req entity.Transaction) error
	FindUserTransaction(userID int) ([]entity.Transaction, error)
	FindAllTransaction() ([]entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) CreateTransaction(transaction entity.Transaction) error {
	err := t.db.Create(&transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionRepository) FindUserTransaction(userID int) ([]entity.Transaction, error) {
	var userTransactions []entity.Transaction
	err := t.db.Preload("Product").Where("user_id = ?", userID).Find(&userTransactions).Error
	if err != nil {
		return nil, err
	}

	return userTransactions, nil
}

func (t *transactionRepository) FindAllTransaction() ([]entity.Transaction, error) {
	var userTransactions []entity.Transaction
	err := t.db.Preload("Product").Find(&userTransactions).Error
	if err != nil {
		return nil, err
	}

	return userTransactions, nil
}
