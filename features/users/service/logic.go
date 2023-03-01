package service

import (
	"clean-arch/features/users"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

// Create implements users.UserServiceInterface
func (userService *userService) Create(input users.UserEntity) error {
	errValidate := userService.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	errInsert := userService.userData.Insert(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

// GetAll implements users.UserServiceInterface
func (userService *userService) GetAll() ([]users.UserEntity, error) {
	userDatas, err := userService.userData.SelectAll()
	return userDatas, err
}

func (userService *userService) GetUser(userId uint) (users.UserEntity, error) {
	userData, err := userService.userData.SelectByUserId(userId)
	return userData, err
}

func (userService *userService) Modify(userId uint, input users.UserEntity) error {
	errUpdate := userService.userData.Update(userId, input)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (userService *userService) Remove(userId uint) error {
	errDelete := userService.userData.Delete(userId)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
