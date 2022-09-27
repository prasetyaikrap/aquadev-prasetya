package handler

import (
	"net/http"
	"strconv"
	"user-management/cases"
	"user-management/entity"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userCases *cases.UserCases
}

func NewUserHandler(userCases *cases.UserCases) *UserHandler {
	return &UserHandler{userCases}
}

func (handler UserHandler) CreateUser(c echo.Context) error {
	req := entity.CreateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := handler.userCases.CreateUser(req)

	if err != nil {
		return err
	}

	return c.JSON(201, user)
}

func (handler UserHandler) GetAllUser(c echo.Context) error{
	users, err := handler.userCases.GetAllUser()
	if err != nil {
		return c.String(http.StatusBadRequest, "Get users failed. Something is wrong.")
	}

	if len(users) <= 0 {
		return c.String(http.StatusNotFound, "Cannot get all users, users is empty or does not exist")
	}

	return c.JSON(200, users)
}

func (handler UserHandler) GetUserById(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	user, err := handler.userCases.GetUserById(intId)

	if err != nil {
		return c.String(http.StatusBadRequest, "Get user by ID failed")
	}

	if user.Id == 0 {
		return c.String(http.StatusNotFound, "Process Failed. User not found or does not exist")
	}

	return c.JSON(200, user)
}