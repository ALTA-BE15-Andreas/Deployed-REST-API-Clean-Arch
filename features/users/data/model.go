package data

import (
	"clean-arch/features/books/data"
	"clean-arch/features/users"

	"gorm.io/gorm"
)

type UserGorm struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Address  string
	Role     string
	Books    []data.BookGorm `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func EntityToGorm(userEntity users.UserEntity) UserGorm {
	return UserGorm{
		Name:     userEntity.Name,
		Email:    userEntity.Email,
		Password: userEntity.Password,
		Address:  userEntity.Address,
		Role:     userEntity.Role,
	}
}

func GormToEntity(userGorm UserGorm) users.UserEntity {
	return users.UserEntity{
		Id:        userGorm.ID,
		Name:      userGorm.Name,
		Email:     userGorm.Email,
		Password:  userGorm.Password,
		Address:   userGorm.Address,
		Role:      userGorm.Role,
		CreatedAt: userGorm.CreatedAt,
		UpdatedAt: userGorm.UpdatedAt,
	}
}

func ListGormToEntity(usersGorm []UserGorm) []users.UserEntity {
	var userEntities []users.UserEntity
	for _, v := range usersGorm {
		userEntities = append(userEntities, GormToEntity(v))
	}
	return userEntities
}
