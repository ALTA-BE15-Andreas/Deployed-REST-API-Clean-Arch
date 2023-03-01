package data

import (
	"clean-arch/features/books"

	"gorm.io/gorm"
)

type BookGorm struct {
	gorm.Model
	Title       string
	UserID      uint
	Author      string
	Publisher   string
	PublishYear string
}

func EntityToGorm(bookEntity books.BookEntity) BookGorm {
	return BookGorm{
		Title:       bookEntity.Title,
		UserID:      bookEntity.UserID,
		Author:      bookEntity.Author,
		Publisher:   bookEntity.Publisher,
		PublishYear: bookEntity.PublishYear,
	}
}

func GormToEntity(bookGorm BookGorm) books.BookEntity {
	return books.BookEntity{
		Id:          bookGorm.ID,
		Title:       bookGorm.Title,
		UserID:      bookGorm.UserID,
		Author:      bookGorm.Author,
		Publisher:   bookGorm.Publisher,
		PublishYear: bookGorm.PublishYear,
		CreatedAt:   bookGorm.CreatedAt,
		UpdatedAt:   bookGorm.UpdatedAt,
	}
}

func ListGormToEntity(booksGorm []BookGorm) []books.BookEntity {
	var bookEntities []books.BookEntity
	for _, v := range booksGorm {
		bookEntities = append(bookEntities, GormToEntity(v))
	}
	return bookEntities
}
