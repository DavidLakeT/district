package main

import (
	"log"

	"district/controller"
	"district/database"
	"district/repository"
	"district/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	database.SetupDatabase(db)
	if dbError != nil {
		log.Fatal(dbError)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository)
	userController := controller.NewUserController(userService)

	app := echo.New()
	app.GET("api/user/:id", userController.Read)
	app.Logger.Fatal(app.Start(":5000"))
}
