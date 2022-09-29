package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type (
	SuccessResponse struct {
		Code int
		Message string
		Data interface{}
	}
)

func HomeAPI(c echo.Context) error{
	return c.JSON(http.StatusOK, SuccessResponse{
			Code: http.StatusOK,
			Message: "API is Active",
			Data: "Home API with Go created by Prasetya ikra priyadi",
		})
}

func main() {
	//init echo instances
	e := echo.New()

	e.GET("/", HomeAPI)

	//Start Server
	e.Logger.Fatal(e.Start(":1323"))
}