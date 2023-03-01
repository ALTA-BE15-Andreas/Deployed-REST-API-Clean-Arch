package delivery

import "clean-arch/features/books"

type BookRequest struct {
	Id          uint   `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	UserID      uint   `json:"user_id" form:"user_id"`
	Author      string `json:"author" form:"author"`
	Publisher   string `json:"publisher" form:"publisher"`
	PublishYear string `json:"publish_year" form:"publish_year"`
}

func requestToEntity(dataRequest BookRequest) books.BookEntity {
	return books.BookEntity{
		Title:       dataRequest.Title,
		UserID:      dataRequest.UserID,
		Author:      dataRequest.Author,
		Publisher:   dataRequest.Publisher,
		PublishYear: dataRequest.PublishYear,
	}
}
