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

	productRepository := repository.NewProductRepository(db)
	reviewRepository := repository.NewReviewRepository(db)
	userRepository := repository.NewUserRepository(db)
	repositoryPool := repositoryPool.NewRepositoryPool(productRepository, reviewRepository, userRepository)

	authService := service.NewAuthService(repositoryPool)
	cartService := service.NewCartService(repositoryPool)
	productService := service.NewProductService(repositoryPool)
	reviewService := service.NewReviewService(repositoryPool)
	userService := service.NewUserService(repositoryPool)
	utilsService := service.NewUtilsService(repositoryPool)
	servicePool := servicePool.NewServicePool(
		authService,
		cartService,
		productService,
		reviewService,
		userService,
		utilsService,
	)

	authController := controller.NewAuthController(servicePool)
	cartController := controller.NewCartController(servicePool)
	productController := controller.NewProductController(servicePool)
	reviewController := controller.NewReviewController(servicePool)
	userController := controller.NewUserController(servicePool)
	utilsController := controller.NewUtilsController(servicePool)
	controllerPool := controllerPool.NewControllerPool(
		authController,
		cartController,
		productController,
		reviewController,
		userController,
		utilsController,
	)

	app := echo.New()
	if err := utilsService.ClearDatabase(); err != nil {
		log.Fatal(err)
	}

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Auth-related endpoints.
	app.POST("api/auth/login", controllerPool.AuthController.LoginUser)

	// Cart-related endpoints.
	app.GET("api/cart", controllerPool.CartController.GetCartInformation)
	app.POST("api/cart", controllerPool.CartController.AddItemToCart)
	app.DELETE("api/cart", controllerPool.CartController.RemoveItemFromCart)

	// Product-related endpoints
	app.GET("api/product", controllerPool.ProductController.GetAllProducts)
	app.GET("api/product/id/:id", controllerPool.ProductController.SearchProductsById)
	app.GET("api/product/name/:name", controllerPool.ProductController.SearchProductsByName)
	app.POST("api/product", controllerPool.ProductController.CreateProduct)
	app.PUT("api/product/id/:id", controllerPool.ProductController.UpdateProduct)
	app.DELETE("api/product/id/:id", controllerPool.ProductController.DeleteProduct)

	// Review-related endpoints
	app.GET("api/review/:id", controllerPool.ReviewController.GetReviewById)
	app.POST("api/review", controllerPool.ReviewController.CreateReview)
	app.DELETE("api/review/:id", controllerPool.ReviewController.DeleteReviewById)

	// User-related endpoints
	app.GET("api/user", controllerPool.UserController.GetAllUsers)
	app.GET("api/user/:id", controllerPool.UserController.GetUserById)
	app.POST("api/user", controllerPool.UserController.CreateUser)
	app.PUT("api/user/:id", controllerPool.UserController.UpdateUser)
	app.DELETE("api/user/:id", controllerPool.UserController.DeleteUser)

	// Database-related endpoints
	app.DELETE("api/database", controllerPool.UtilsController.ClearDatabase)

	app.Logger.Fatal(app.Start(":5000"))
}
