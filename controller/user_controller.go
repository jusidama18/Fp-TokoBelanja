package controller

import (
	"TokoBelanja/helper"
	"TokoBelanja/model/input"
	"TokoBelanja/model/response"
	"TokoBelanja/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

// @Summary Register New User
// @Description Register New User by Data Provided
// @Tags Users
// @Accept json
// @Produce json
// @Param data body input.UserRegisterInput true "Register User"
// @Success 200 {object} helper.Response{data=response.UserRegisterResponse}
// @Router /users/register [post]
func (h *userController) RegisterUser(c *gin.Context) {
	var input input.UserRegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	userData, err := h.userService.RegisterUser(input)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	userResponse := response.UserRegisterResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		Password:  userData.Password,
		Balance:   userData.Balance,
		CreatedAt: userData.CreatedAt,
	}

	helper.Success(c, http.StatusCreated, "created", userResponse)
}

// @Summary Login Account
// @Description Login Account by Data Provided
// @Tags Users
// @Accept json
// @Produce json
// @Param data body input.UserLoginInput true "Login Account"
// @Success 200 {object} helper.Response{data=response.UserLoginResponse}
// @Router /users/login [post]
func (h *userController) LoginUser(c *gin.Context) {
	var input input.UserLoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	token, err := h.userService.LoginUser(input)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	userResponse := response.UserLoginResponse{
		Token: token,
	}

	helper.Success(c, http.StatusOK, "ok", userResponse)
}

// @Summary Patch User's Topup
// @Description Patch User's Topup by Data Provided
// @Tags Users
// @Accept json
// @Produce json
// @Param data body input.UserPatchTopUpInput true "Patch User's Topup"
// @Success 200 {object} helper.Response{data=response.UserPatchTopUpResponse}
// @Router /users/topup [patch]
func (h *userController) PatchTopUpUser(c *gin.Context) {
	var input input.UserPatchTopUpInput

	id_user := c.MustGet("currentUser").(int)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	userData, err := h.userService.TopUpUser(id_user, input)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	message := "Your balance has been succcessfully updated to Rp " + strconv.Itoa(userData.Balance)
	userResponse := response.UserPatchTopUpResponse{
		Message: message,
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			userResponse,
		),
	)
}

// @Summary Register New Admin
// @Description Register New Admin by Data Provided
// @Tags Users
// @Accept json
// @Produce json
// @Param data body input.UserRegisterInput true "Register Admin"
// @Success 200 {object} helper.Response{data=response.UserRegisterResponse}
// @Router /users/admin [post]
func (h *userController) RegisterAdmin(c *gin.Context) {
	var input input.UserRegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	userData, err := h.userService.RegisterAdmin(input)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	userResponse := response.UserRegisterResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		Password:  userData.Password,
		Balance:   userData.Balance,
		CreatedAt: userData.CreatedAt,
	}
	helper.Success(c, http.StatusCreated, "created", userResponse)
}
