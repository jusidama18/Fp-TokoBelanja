package controller

import (
	"TokoBelanja/helper"
	"TokoBelanja/model/input"
	"TokoBelanja/model/response"
	"TokoBelanja/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *categoryController {
	return &categoryController{categoryService}
}

// @Summary Create Category
// @Description Create Category by Data Provided
// @Tags Categories
// @Accept json
// @Produce json
// @Param data body input.CategoryCreateInput true "Create Category"
// @Success 200 {object} helper.Response{data=response.CategoryCreateResponse}
// @Router /categories [post]
func (h *categoryController) CreateCategory(c *gin.Context) {
	var input input.CategoryCreateInput

	role_user := c.MustGet("roleUser").(string)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	categoryData, err := h.categoryService.CreateCategory(role_user, input)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	categoryResponse := response.CategoryCreateResponse{
		ID:                categoryData.ID,
		Type:              categoryData.Type,
		SoldProductAmount: categoryData.SoldProductAmount,
		CreatedAt:         categoryData.CreatedAt,
	}

	helper.Success(c, http.StatusCreated, "created", categoryResponse)
}

// @Summary Get All Category with Product
// @Description Get All Category with Product
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response{data=[]response.CategoryGetResponse}
// @Router /categories [get]
func (h *categoryController) GetAllCategories(c *gin.Context) {
	var (
		allProducts   []response.CategoryProduct
		allCategories []response.CategoryGetResponse
	)

	role_user := c.MustGet("roleUser").(string)

	categoryData, err := h.categoryService.GetAllCategories(role_user)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	for _, dataCategory := range categoryData {
		productsData, err := h.categoryService.GetProductsByCategoryID(dataCategory.ID)
		if err != nil {
			errors := helper.GetErrorData(err)
			helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
			return
		}
		for _, dataProduct := range productsData {
			allProductsTmp := response.CategoryProduct{
				ID:        dataProduct.ID,
				Title:     dataProduct.Title,
				Price:     dataProduct.Price,
				Stock:     dataProduct.Stock,
				CreatedAt: dataProduct.CreatedAt,
				UpdatedAt: dataProduct.UpdatedAt,
			}
			allProducts = append(allProducts, allProductsTmp)
		}
		allCategoriesTmp := response.CategoryGetResponse{
			ID:                dataCategory.ID,
			Type:              dataCategory.Type,
			SoldProductAmount: dataCategory.SoldProductAmount,
			CreatedAt:         dataCategory.CreatedAt,
			UpdatedAt:         dataCategory.UpdatedAt,
			Product:           allProducts,
		}

		allCategories = append(allCategories, allCategoriesTmp)
		allProducts = []response.CategoryProduct{}
	}

	helper.Success(c, http.StatusOK, "ok", allCategories)
}

// @Summary Patch Category
// @Description Patch Category by Data Provided
// @Tags Categories
// @Accept json
// @Produce json
// @Param data body input.CategoryPatchInput true "Patch Category"
// @Param id path int true "Category ID"
// @Success 200 {object} helper.Response{data=response.CategoryPatchResponse}
// @Router /categories/{id} [patch]
func (h *categoryController) PatchCategory(c *gin.Context) {
	var (
		inputBody input.CategoryPatchInput
		uri       input.CategoryIdUri
	)

	role_user := c.MustGet("roleUser").(string)

	err := c.ShouldBindJSON(&inputBody)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	err = c.ShouldBindUri(&uri)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	categoryData, err := h.categoryService.PatchCategory(role_user, uri.ID, inputBody)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	categoryResponse := response.CategoryPatchResponse{
		ID:                categoryData.ID,
		Type:              categoryData.Type,
		SoldProductAmount: categoryData.SoldProductAmount,
		UpdatedAt:         categoryData.UpdatedAt,
	}

	helper.Success(c, http.StatusOK, "ok", categoryResponse)
}

// @Summary Delete Category
// @Description Delete Category by Data Provided
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Delete Category"
// @Success 200 {object} helper.Response{data=string}
// @Router /categories/{id} [delete]
func (h *categoryController) DeleteCategory(c *gin.Context) {
	var uri input.CategoryIdUri

	role_user := c.MustGet("roleUser").(string)

	err := c.ShouldBindUri(&uri)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	err = h.categoryService.DeleteCategory(role_user, uri.ID)
	if err != nil {
		errors := helper.GetErrorData(err)
		helper.Error(c, http.StatusUnprocessableEntity, "failed", errors)
		return
	}

	categoryResponse := response.CategoryDeleteResponse{
		Message: "Category has been successfully deleted",
	}

	helper.Success(c, http.StatusOK, "ok", categoryResponse)
}
