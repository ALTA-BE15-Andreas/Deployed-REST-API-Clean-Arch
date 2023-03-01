package router

import (
	_userData "clean-arch/features/users/data"
	_userHandler "clean-arch/features/users/delivery"
	_userService "clean-arch/features/users/service"

	_bookData "clean-arch/features/books/data"
	_bookHandler "clean-arch/features/books/delivery"
	_bookService "clean-arch/features/books/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initUserRouter (db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	e.GET("/users", userHandlerAPI.GetAllUser)
	e.POST("/users", userHandlerAPI.Register)
	e.PUT("/users/:id", userHandlerAPI.UpdateAccount)
	e.DELETE("/users/:id", userHandlerAPI.RemoveAccount)
}

func initBookRouter (db *gorm.DB, e *echo.Echo) {
	bookData := _bookData.New(db)
	bookService := _bookService.New(bookData)
	bookHandlerAPI := _bookHandler.New(bookService)

	e.GET("/books", bookHandlerAPI.GetAllBook)
	e.POST("/books", bookHandlerAPI.CreateBook)
	e.PUT("/books/:id", bookHandlerAPI.UpdateBook)
	e.DELETE("/books/:id", bookHandlerAPI.RemoveBook)
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	initUserRouter(db, e)
	initBookRouter(db, e)
}
