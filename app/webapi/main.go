package main

import (
	"log"

	"district/controller"
	"district/database"
	"district/repository"
	repositoryPool "district/repository/handler"
	"district/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	envError := godotenv.Load("config.env")
	if envError != nil {
		log.Fatal(envError)
	}

	db, dbError := database.ConnectDatabase()
	if dbError != nil {
		log.Fatal(dbError)
	}

	dbError = database.SetupDatabase(db)
	if dbError != nil {
		log.Fatal(dbError)
	}

	productRepository := repository.NewProductRepository(db)
	reviewRepository := repository.NewReviewRepository(db)
	userRepository := repository.NewUserRepository(db)
	repositoryPool := repositoryPool.NewRepositoryPool(productRepository, reviewRepository, userRepository)

	authService := service.NewAuthService(repositoryPool)
	productService := service.NewProductService(repositoryPool)
	reviewService := service.NewReviewService(repositoryPool)
	userService := service.NewUserService(repositoryPool)
	authController := controller.NewAuthController(authService)
	productController := controller.NewProductController(productService)
	reviewController := controller.NewReviewController(reviewService)
	userController := controller.NewUserController(userService)

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Auth-related endpoints.
	app.POST("api/auth/login", authController.LoginUser)

	// Product-related endpoints
	app.GET("api/product", productController.GetAllProducts)
	app.GET("api/product/id/:id", productController.SearchProductsById)
	app.GET("api/product/name/:name", productController.SearchProductsByName)
	app.POST("api/product", productController.CreateProduct)

	// Review-related endpoints
	app.GET("api/review/:id", reviewController.GetReviewById)
	app.POST("api/review", reviewController.CreateReview)

	// User-related endpoints
	app.GET("api/user/:id", userController.GetUserById)
	app.POST("api/user", userController.CreateUser)
	app.PUT("api/user/:id", userController.UpdateUser)
	app.DELETE("api/user/:id", userController.DeleteUser)

	app.Logger.Fatal(app.Start(":5000"))
}
