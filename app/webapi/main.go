package main

import (
	"log"

	"district/controller"
	controllerPool "district/controller/handler"
	"district/database"
	"district/repository"
	repositoryPool "district/repository/handler"
	"district/service"
	servicePool "district/service/handler"

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
	servicePool := servicePool.NewServicePool(authService, productService, reviewService, userService)

	authController := controller.NewAuthController(servicePool)
	productController := controller.NewProductController(servicePool)
	reviewController := controller.NewReviewController(servicePool)
	userController := controller.NewUserController(servicePool)
	controllerPool := controllerPool.NewControllerPool(authController, productController, reviewController, userController)

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Auth-related endpoints.
	app.POST("api/auth/login", controllerPool.AuthController.LoginUser)

	// Product-related endpoints
	app.GET("api/product", controllerPool.ProductController.GetAllProducts)
	app.GET("api/product/id/:id", controllerPool.ProductController.SearchProductsById)
	app.GET("api/product/name/:name", controllerPool.ProductController.SearchProductsByName)
	app.POST("api/product", controllerPool.ProductController.CreateProduct)

	// Review-related endpoints
	app.GET("api/review/:id", controllerPool.ReviewController.GetReviewById)
	app.POST("api/review", controllerPool.ReviewController.CreateReview)

	// User-related endpoints
	app.GET("api/user/:id", controllerPool.UserController.GetUserById)
	app.POST("api/user", controllerPool.UserController.CreateUser)
	app.PUT("api/user/:id", controllerPool.UserController.UpdateUser)
	app.DELETE("api/user/:id", controllerPool.UserController.DeleteUser)

	app.Logger.Fatal(app.Start(":5000"))
}
