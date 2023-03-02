package data

import (
	"clean-arch/features/users"
	"errors"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Create implements users.UserData
func (repo *userQuery) Insert(input users.UserEntity) error {
	userGorm := EntityToGorm(input)
	txInsert := repo.db.Create(&userGorm)
	if txInsert.Error != nil {
		return txInsert.Error
	}

	if txInsert.RowsAffected == 0 {
		return errors.New("insert error, row affected = 0")
	}
	return nil
}

// GetAll implements users.UserData
func (repo *userQuery) SelectAll() ([]users.UserEntity, error) {
	usersGorm := []UserGorm{}
	txSelect := repo.db.Find(&usersGorm)
	if txSelect.Error != nil {
		return nil, txSelect.Error
	}

	userEntities := ListGormToEntity(usersGorm)
	return userEntities, nil
}

func (repo *userQuery) SelectByUserId(userId uint) (users.UserEntity, error) {
	userGorm := UserGorm{}
	txSelect := repo.db.Find(&userGorm, userId)

	userEntity := GormToEntity(userGorm)
	if txSelect.Error != nil {
		return users.UserEntity{}, errors.New("error select user")
	}
	return userEntity, nil
}

func (repo *userQuery) Update(userId uint, input users.UserEntity) error {
	selectedUserEntity, err := repo.SelectByUserId(userId)
	if err != nil {
		return err
	}

	if input.Name != "" {
		selectedUserEntity.Name = input.Name
	}
	if input.Email != "" {
		selectedUserEntity.Email = input.Email
	}
	if input.Address != "" {
		selectedUserEntity.Address = input.Address
	}
	if input.Role != "" {
		selectedUserEntity.Role = input.Role
	}
	selectedUserGorm := EntityToGorm(selectedUserEntity)
	selectedUserGorm.ID = userId
	txUpdate := repo.db.Model(&selectedUserGorm).Where("id = ?", userId).Save(&selectedUserGorm)
	if txUpdate.Error != nil {
		return txUpdate.Error
	}

	if txUpdate.RowsAffected == 0 {
		return errors.New("update error, row affected = 0")
	}
	return nil
}

// GetAll implements users.UserData
func (repo *userQuery) Delete(userId uint) error {
	txDelete := repo.db.Delete(&UserGorm{}, userId)
	if txDelete.Error != nil {
		return errors.New("delete error, row affected = 0")
	}
	return nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &userQuery{
		db: db,
	}
}
