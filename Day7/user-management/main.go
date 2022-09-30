package main

import (
	"user-management/cases"
	"user-management/config"
	"user-management/handler"
	"user-management/repository"
	"user-management/route"

	echo "github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	config.Migrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := cases.NewUserCases(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	route.Routes(e, userHandler)
	
	e.Logger.Fatal(e.Start(":4000"))
}