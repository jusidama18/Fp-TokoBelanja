package main

import (
	"TokoBelanja/config"
	"TokoBelanja/controller"
	"TokoBelanja/middleware"
	"TokoBelanja/repository"
	"TokoBelanja/service"
	"os"

	"github.com/gin-gonic/gin"

	_ "TokoBelanja/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title TokoBelanja-API
// @version 1.0
// @description This is a API webservice to manage TokoBelanja API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email hacktiv@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
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

	docs := router.Group("/docs")
	docs.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
