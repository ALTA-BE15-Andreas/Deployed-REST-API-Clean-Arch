package books

import "time"

type BookEntity struct {
	Id          uint
	Title       string
	UserID      uint
	Author      string
	Publisher   string
	PublishYear string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type BookDataInterface interface {
	SelectAll() ([]BookEntity, error)
	SelectByBookId(bookId uint) (BookEntity, error)
	Insert(input BookEntity) error
	Update(bookId uint, input BookEntity) error
	Delete(bookId uint) error
}

type BookServiceInterface interface {
	GetAll() ([]BookEntity, error)
	GetBook(bookId uint) (BookEntity, error)
	Create(input BookEntity) error
	Modify(bookId uint, input BookEntity) error
	Remove(bookId uint) error
}
