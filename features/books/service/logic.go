package service

import (
	"clean-arch/features/books"
)

type bookService struct {
	bookData books.BookDataInterface
}

func (bookService *bookService) Create(input books.BookEntity) error {
	errInsert := bookService.bookData.Insert(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func (bookService *bookService) GetAll() ([]books.BookEntity, error) {
	bookDatas, err := bookService.bookData.SelectAll()
	return bookDatas, err
}

func (bookService *bookService) GetBook(bookId uint) (books.BookEntity, error) {
	bookData, err := bookService.bookData.SelectByBookId(bookId)
	return bookData, err
}

func (bookService *bookService) Modify(bookId uint, input books.BookEntity) error {
	errUpdate := bookService.bookData.Update(bookId, input)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (bookService *bookService) Remove(bookId uint) error {
	errDelete := bookService.bookData.Delete(bookId)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

func New(repo books.BookDataInterface) books.BookServiceInterface {
	return &bookService{
		bookData: repo,
	}
}
