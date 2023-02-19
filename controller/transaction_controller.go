package controller

import (
	"TokoBelanja/helper"
	"TokoBelanja/model/input"
	"TokoBelanja/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionController struct {
	transService service.TransactionService
}

func NewTransactionController(transService service.TransactionService) *transactionController {
	return &transactionController{
		transService: transService,
	}
}

// @Summary Create Transaction
// @Description Create Transaction by Data Provided
// @Tags Transactions
// @Accept json
// @Produce json
// @Param data body input.CreateTransactionRequest true "Create Transaction"
// @Success 200 {object} helper.Response{data=response.TransactionBillResponse}
// @Router /transactions [post]
func (t *transactionController) CreateTransaction(c *gin.Context) {
	var req input.CreateTransactionRequest

	userID := c.MustGet("currentUser").(int)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.Error(c, http.StatusUnprocessableEntity, "BAD_REQUEST", err.Error())
		return
	}

	err = c.ShouldBindUri(&req)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusBadRequest, "BAD_REQUEST", errors)
		return
	}

	resp, err := t.transService.CreateTransaction(userID, req)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
		return
	}
	helper.Success(c, http.StatusCreated, "You have successfully purchased the product", resp)
}

// @Summary Get My Transaction
// @Description Get My Transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response{data=[]response.MyTransactionResponse}
// @Router /transactions/my-transactions [get]
func (t *transactionController) FindMyTransactions(c *gin.Context) {
	userID := c.MustGet("currentUser").(int)
	resp, err := t.transService.FindMyTransaction(userID)
	if err != nil {
		helper.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", err.Error())
		return
	}
	helper.Success(c, http.StatusCreated, "successfully get your transaction history", resp)
}

// @Summary Get All User Transaction
// @Description Get All User Transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response{data=[]response.UserTransactionResponse}
// @Router /transactions/user-transactions [get]
func (t *transactionController) FindUserTransaction(c *gin.Context) {
	userID := c.MustGet("currentUser").(int)
	userRole := c.MustGet("roleUser").(string)

	if userRole != "admin" {
		helper.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", gin.H{"error": "user not authorized"})
		return
	}

	resp, err := t.transService.FindUserTransaction(userID)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
		return
	}

	helper.Success(c, http.StatusCreated, "successfully get all transaction history", resp)
}
