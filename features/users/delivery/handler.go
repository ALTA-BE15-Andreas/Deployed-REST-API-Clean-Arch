package delivery

import (
	"clean-arch/features/users"
	"clean-arch/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func New(srv users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: srv,
	}
}

func (delivery *UserHandler) Register(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	userEntity := requestToEntity(userInput)
	err := delivery.userService.Create(userEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error: "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil insert data user"))
}

func (delivery *UserHandler) GetAllUser(c echo.Context) error {
	data, err := delivery.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
	}
	dataResponse := listEntityToResponse(data)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil read user data", dataResponse))
}

func (delivery *UserHandler) UpdateAccount(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	userId, errCasting := strconv.Atoi(c.Param("id"))
	if errCasting != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invaild param"))
	}

	userEntity := requestToEntity(userInput)
	err := delivery.userService.Modify(uint(userId), userEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error: "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil update data user"))
}

func (delivery *UserHandler) RemoveAccount(c echo.Context) error {
	userId, errCasting := strconv.Atoi(c.Param("id"))
	if errCasting != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invaild param"))
	}
	err := delivery.userService.Remove(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil delete data user"))
}
