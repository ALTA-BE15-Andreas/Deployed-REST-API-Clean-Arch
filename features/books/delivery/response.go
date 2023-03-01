package delivery

import (
	"clean-arch/features/books"
)

type BookResponse struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	UserID      uint   `json:"user_id"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	PublishYear string `json:"publish_year"`
}

func entityToResponse(bookEntity books.BookEntity) BookResponse {
	return BookResponse{
		Id:          bookEntity.Id,
		Title:       bookEntity.Title,
		UserID:      bookEntity.UserID,
		Author:      bookEntity.Author,
		Publisher:   bookEntity.Publisher,
		PublishYear: bookEntity.PublishYear,
	}
}

func listEntityToResponse(bookEntities []books.BookEntity) []BookResponse {
	var bookResponse []BookResponse
	for _, v := range bookEntities {
		bookResponse = append(bookResponse, entityToResponse(v))
	}
	return bookResponse
}
