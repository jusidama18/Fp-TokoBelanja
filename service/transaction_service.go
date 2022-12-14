package service

import (
	"TokoBelanja/model/entity"
	"TokoBelanja/model/input"
	"TokoBelanja/model/response"
	"TokoBelanja/repository"
	"fmt"
)

type TransactionService interface {
	CreateTransaction(userId int, req input.CreateTransactionRequest) (*response.TransactionBillResponse, error)
	FindMyTransaction(userID int) ([]response.MyTransactionResponse, error)
	FindUserTransaction(userID int) ([]response.UserTransactionResponse, error)
}

type transactionService struct {
	transRepo   repository.TransactionRepository
	productRepo repository.ProductsRepository
	userRepo    repository.UserRepository
}

func NewTransactionService(transRepo repository.TransactionRepository, productRepo repository.ProductsRepository, userRepo repository.UserRepository) *transactionService {
	return &transactionService{
		transRepo:   transRepo,
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (t *transactionService) CreateTransaction(userID int, req input.CreateTransactionRequest) (*response.TransactionBillResponse, error) {
	productID := req.ProductID
	product, err := t.productRepo.FindById(productID)
	if err != nil {
		fmt.Println("prod")
		return nil, err
	}

	if req.Quantity > product.Stock {
		fmt.Println("stock")
		return nil, fmt.Errorf("%s's stock is not enough. Stock: %v", product.Title, product.Stock)
	}

	user, err := t.userRepo.FindById(userID)
	fmt.Println(user)
	if err != nil {
		fmt.Println("user")
		return nil, err
	}

	if user == *new(entity.User) {
		return nil, fmt.Errorf("User with id %d not found", userID)
	}

	totalCost := product.Price * req.Quantity

	if user.Balance < totalCost {
		fmt.Println("balance")
		return nil, fmt.Errorf("Balance is not enough. Balance: %v", user.Balance)
	}

	transaction := entity.Transaction{
		UserID:     userID,
		TotalPrice: totalCost,
		ProductID:  productID,
		Quantity:   req.Quantity,
	}

	err = t.transRepo.CreateTransaction(transaction)
	if err != nil {
		fmt.Println("trans")
		return nil, fmt.Errorf("Failed to create transaction: %v", err)
	}

	newStock := product.Stock - req.Quantity
	product.Stock = newStock
	_, err = t.productRepo.Update(productID, product)
	if err != nil {
		fmt.Println("update prod")
		return nil, fmt.Errorf("Failed to update new product stock, %v", err)
	}

	newBalance := user.Balance - totalCost
	user.Balance = newBalance
	_, err = t.userRepo.Update(userID, user)
	if err != nil {
		return nil, err
	}

	transactionBill := &response.TransactionBillResponse{
		TotalPrice:   totalCost,
		Quantity:     req.Quantity,
		ProductTitle: product.Title,
	}

	return transactionBill, nil
}

func (t *transactionService) FindMyTransaction(userID int) ([]response.MyTransactionResponse, error) {
	user, err := t.userRepo.FindById(userID)
	if user == *new(entity.User) {
		return nil, fmt.Errorf("User with id %d not found", userID)
	}

	userTransactions, err := t.transRepo.FindUserTransaction(userID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get user transactions: %v", err)
	}

	parsedTrans := parseMultipleMyTransaction(userTransactions)

	return parsedTrans, nil
}

func (t *transactionService) FindUserTransaction(userID int) ([]response.UserTransactionResponse, error) {
	user, err := t.userRepo.FindById(userID)
	if user == *new(entity.User) {
		return nil, fmt.Errorf("User with id %d not found", userID)
	}

	userTransactions, err := t.transRepo.FindAllTransaction()
	if err != nil {
		return nil, fmt.Errorf("Failed to get all transactions: %v", err)
	}

	parsedTrans := parseMultiAllTransactions(userTransactions)

	return parsedTrans, nil
}

func parseTransactionProduct(trans entity.Transaction) response.TransactionProduct {
	return response.TransactionProduct{
		ID:         trans.Product.ID,
		Title:      trans.Product.Title,
		Price:      trans.Product.Price,
		Stock:      trans.Product.Stock,
		CategoryID: trans.Product.CategoryID,
		CreatedAt:  trans.CreatedAt,
		UpdatedAt:  trans.UpdatedAt,
	}
}

func parseTransactionUser(trans entity.Transaction) response.TransactionUser {
	return response.TransactionUser{
		ID:        trans.User.ID,
		Email:     trans.User.Email,
		FullName:  trans.User.FullName,
		Balance:   trans.User.Balance,
		CreatedAt: trans.User.CreatedAt,
		UpdatedAt: trans.User.UpdatedAt,
	}
}

func parseAllTransaction(trans entity.Transaction) response.UserTransactionResponse {
	return response.UserTransactionResponse{
		ID:         trans.ID,
		ProductID:  trans.ProductID,
		UserID:     trans.UserID,
		Quantity:   trans.Quantity,
		TotalPrice: trans.TotalPrice,
		Product:    parseTransactionProduct(trans),
		User:       parseTransactionUser(trans),
	}
}

func parseMultiAllTransactions(trans []entity.Transaction) []response.UserTransactionResponse {
	var resp []response.UserTransactionResponse
	for _, t := range trans {
		resp = append(resp, parseAllTransaction(t))
	}
	return resp
}

func parseMyTransaction(trans entity.Transaction) response.MyTransactionResponse {
	return response.MyTransactionResponse{
		ID:         trans.ID,
		ProductID:  trans.ProductID,
		UserID:     trans.UserID,
		Quantity:   trans.Quantity,
		TotalPrice: trans.TotalPrice,
		Product:    parseTransactionProduct(trans),
	}
}

func parseMultipleMyTransaction(trans []entity.Transaction) []response.MyTransactionResponse {
	var parsedTrans []response.MyTransactionResponse
	for _, t := range trans {
		parsedTrans = append(parsedTrans, parseMyTransaction(t))
	}

	return parsedTrans
}
