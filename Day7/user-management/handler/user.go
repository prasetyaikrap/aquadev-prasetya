package handler

import (
	"net/http"
	"strconv"
	"user-management/cases"
	"user-management/entity"

	echo "github.com/labstack/echo/v4"
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
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Create user failed",
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.SuccessResponse{
		Code: http.StatusCreated,
		Message: "User created successfully",
		Data: user,
	})
}

func (handler UserHandler) GetAllUser(c echo.Context) error{
	users, err := handler.userCases.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Get users Failed",
			Error: err.Error(),
		})
	}

	if len(users) <= 0 {
		return c.JSON(http.StatusNotFound, entity.ErrorResponse{
			Code: http.StatusNotFound,
			Message: "Get users Failed",
			Error: "Users not found",
		})
	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code: http.StatusOK,
		Message: "Users found",
		Data: users,
	})
}

func (handler UserHandler) GetUserById(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	user, err := handler.userCases.GetUserById(intId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Get user Failed",
			Error: err.Error(),
		})
	}

	if user.Id == 0 {
		return c.JSON(http.StatusNotFound, entity.ErrorResponse{
			Code: http.StatusNotFound,
			Message: "Get user Failed.",
			Error: "User not found",
		})
	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code: http.StatusOK,
		Message: "User found",
		Data: user,
	})
}

func (handler UserHandler) UpdateUser(c echo.Context) error {
	req := entity.UpdateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Request Body",
			Error: err.Error(),
		})
	}

	intId, _ := strconv.Atoi(c.Param("id"))
	user, err := handler.userCases.UpdateUser(req, intId)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Update user failed",
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code: http.StatusOK,
		Message: "User updated successfully",
		Data: user,
	})
}

func (handler UserHandler) DeleteUser(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	err := handler.userCases.DeleteUser(intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Delete user failed",
			Error: err.Error(),
		}))
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code: http.StatusOK,
		Message: "User deleted successfully",
	})
}