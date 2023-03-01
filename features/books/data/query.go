package data

import (
	"clean-arch/features/books"
	"errors"

	"gorm.io/gorm"
)

type bookQuery struct {
	db *gorm.DB
}

// Create implements users.UserData
func (repo *bookQuery) Insert(input books.BookEntity) error {
	bookGorm := EntityToGorm(input)
	txInsert := repo.db.Create(&bookGorm)
	if txInsert.Error != nil {
		return txInsert.Error
	}

	if txInsert.RowsAffected == 0 {
		return errors.New("insert error, row affected = 0")
	}
	return nil
}

// GetAll implements users.UserData
func (repo *bookQuery) SelectAll() ([]books.BookEntity, error) {
	booksGorm := []BookGorm{}
	txSelect := repo.db.Find(&booksGorm)
	if txSelect.Error != nil {
		return nil, txSelect.Error
	}

	userEntities := ListGormToEntity(booksGorm)
	return userEntities, nil
}

func (repo *bookQuery) SelectByBookId(bookId uint) (books.BookEntity, error) {
	bookGorm := BookGorm{}
	txSelect := repo.db.Find(&bookGorm, bookId)

	bookEntity := GormToEntity(bookGorm)
	if txSelect.Error != nil {
		return books.BookEntity{}, txSelect.Error
	}
	return bookEntity, nil
}

func (repo *bookQuery) Update(bookId uint, input books.BookEntity) error {
	selectedBookEntity, err := repo.SelectByBookId(bookId)
	if err != nil {
		return err
	}

	if input.Title != "" {
		selectedBookEntity.Title = input.Title
	}
	if input.UserID != 0 {
		selectedBookEntity.UserID = input.UserID
	}
	if input.Author != "" {
		selectedBookEntity.Author = input.Author
	}
	if input.Publisher != "" {
		selectedBookEntity.Publisher = input.Publisher
	}
	if input.PublishYear != "" {
		selectedBookEntity.PublishYear = input.PublishYear
	}
	selectedBookGorm := EntityToGorm(selectedBookEntity)
	selectedBookGorm.ID = bookId
	txUpdate := repo.db.Model(&selectedBookGorm).Where("id = ?", bookId).Save(&selectedBookGorm)
	if txUpdate.Error != nil {
		return txUpdate.Error
	}

	if txUpdate.RowsAffected == 0 {
		return errors.New("update error, row affected = 0")
	}
	return nil
}

// GetAll implements users.UserData
func (repo *bookQuery) Delete(bookId uint) error {
	txDelete := repo.db.Delete(&BookGorm{}, bookId)
	if txDelete.Error != nil {
		return txDelete.Error
	}
	return nil
}

func New(db *gorm.DB) books.BookDataInterface {
	return &bookQuery{
		db: db,
	}
}

