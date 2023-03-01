package delivery

import (
	"clean-arch/features/books"
	"clean-arch/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookService books.BookServiceInterface
}

func New(srv books.BookServiceInterface) *BookHandler {
	return &BookHandler{
		bookService: srv,
	}
}

func (delivery *BookHandler) CreateBook(c echo.Context) error {
	bookInput := BookRequest{}
	errBind := c.Bind(&bookInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	bookEntity := requestToEntity(bookInput)
	err := delivery.bookService.Create(bookEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error: "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil insert data book"))
}

func (delivery *BookHandler) GetAllBook(c echo.Context) error {
	data, err := delivery.bookService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
	}
	dataResponse := listEntityToResponse(data)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("succesfulyy read book data", dataResponse))
}

func (delivery *BookHandler) UpdateBook(c echo.Context) error {
	bookInput := BookRequest{}
	errBind := c.Bind(&bookInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	bookId, errCasting := strconv.Atoi(c.Param("id"))
	if errCasting != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invaild param"))
	}

	bookEntity := requestToEntity(bookInput)
	err := delivery.bookService.Modify(uint(bookId), bookEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error: "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil update data book"))
}

func (delivery *BookHandler) RemoveBook(c echo.Context) error {
	bookId, errCasting := strconv.Atoi(c.Param("id"))
	if errCasting != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invaild param"))
	}
	err := delivery.bookService.Remove(uint(bookId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil delete data book"))
}