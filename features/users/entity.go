package users

import "time"

type UserEntity struct {
	Id        uint
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string
	Address   string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDataInterface interface {
	SelectAll() ([]UserEntity, error)
	SelectByUserId(userId uint) (UserEntity, error)
	Insert(input UserEntity) error
	Update(userId uint, input UserEntity) error
	Delete(userId uint) error
}

type UserServiceInterface interface {
	GetAll() ([]UserEntity, error)
	GetUser(userId uint) (UserEntity, error)
	Create(input UserEntity) error
	Modify(userId uint, input UserEntity) error
	Remove(userId uint) error
}
