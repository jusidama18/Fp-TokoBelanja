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

func (t *transactionController) CreateTransaction(c *gin.Context) {
	var req input.CreateTransactionRequest

	userID := c.MustGet("currentUser").(int)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.NewErrorResponse(http.StatusUnprocessableEntity, "BAD_REQUEST", err.Error()))
		return
	}

	err = c.ShouldBindUri(&req)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.NewErrorResponse(http.StatusBadRequest, "BAD_REQUEST", errors))
		return
	}

	resp, err := t.transService.CreateTransaction(userID, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.NewErrorResponse(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error()))
		return
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(http.StatusCreated, "You have successfully purchased the product", resp),
	)
}

func (t *transactionController) FindMyTransactions(c *gin.Context) {
	userID := c.MustGet("currentUser").(int)
	resp, err := t.transService.FindMyTransaction(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helper.NewErrorResponse(http.StatusInternalServerError, "UNAUTHORIZED", err.Error()))
		return
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(http.StatusCreated, "successfully get your transaction history", resp),
	)
}

func (t *transactionController) FindUserTransaction(c *gin.Context) {
	userID := c.MustGet("currentUser").(int)
	userRole := c.MustGet("roleUser").(string)

	if userRole != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helper.NewErrorResponse(http.StatusInternalServerError, "UNAUTHORIZED", gin.H{
			"error": "user not authorized",
		}))
		return
	}

	resp, err := t.transService.FindUserTransaction(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.NewErrorResponse(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error()))
		return
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(http.StatusCreated, "successfully get all transaction history", resp),
	)
}
