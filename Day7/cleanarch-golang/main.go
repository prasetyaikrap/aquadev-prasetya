package main

import (
	"cleanarch-golang/config"
	"cleanarch-golang/handler"
	"cleanarch-golang/repository"
	"cleanarch-golang/route"
	"cleanarch-golang/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	config.Migrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	route.Routes(e, userHandler)
	
	e.Logger.Fatal(e.Start(":4000"))
}