package main

import (
	"log"

	"district/controller"
	"district/database"
	"district/repository"
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

	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository)
	productService := service.NewProductService(productRepository)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)
	productController := controller.NewProductController(productService)

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	app.GET("api/user/:id", userController.GetUserInformation)
	app.POST("api/user", userController.CreateUser)
	app.POST("api/auth/login", authController.LoginUser)

	app.GET("api/product/name/:name", productController.SearchProductsByName)
	app.GET("api/product/id/:id", productController.SearchProductsById)
	app.GET("api/product", productController.GetAllProducts)
	app.POST("api/product", productController.CreateProduct)

	app.Logger.Fatal(app.Start(":5000"))
}
