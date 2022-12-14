package main

import (
	"TokoBelanja/config"
	"TokoBelanja/controller"
	"TokoBelanja/middleware"
	"TokoBelanja/repository"
	"TokoBelanja/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.StartDB()
	db := config.GetDBConnection()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo, productRepository, userRepository)
	transactionController := controller.NewTransactionController(transactionService)

	// Users
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.LoginUser)
	router.PATCH("/users/topup", middleware.AuthMiddleware, userController.PatchTopUpUser)
	// Create Admin
	router.POST("/users/admin", userController.RegisterAdmin)

	// Categories
	categoryGroup := router.Group("/categories")
	categoryGroup.POST("/", middleware.AuthMiddleware, categoryController.CreateCategory)
	categoryGroup.GET("/", middleware.AuthMiddleware, categoryController.GetAllCategories)
	categoryGroup.PATCH("/:id", middleware.AuthMiddleware, categoryController.PatchCategory)
	categoryGroup.DELETE("/:id", middleware.AuthMiddleware, categoryController.DeleteCategory)

	// Product
	productGroup := router.Group("/products")
	productGroup.POST("/", middleware.AuthMiddleware, productController.Post)
	productGroup.GET("/", middleware.AuthMiddleware, productController.Get)
	productGroup.PUT("/:id", middleware.AuthMiddleware, productController.Put)
	productGroup.DELETE("/:id", middleware.AuthMiddleware, productController.Delete)

	transGroup := router.Group("/transactions")
	transGroup.POST("/", middleware.AuthMiddleware, transactionController.CreateTransaction)
	transGroup.GET("/my-transactions", middleware.AuthMiddleware, transactionController.FindMyTransactions)
	transGroup.GET("/user-transactions", middleware.AuthMiddleware, transactionController.FindUserTransaction)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
