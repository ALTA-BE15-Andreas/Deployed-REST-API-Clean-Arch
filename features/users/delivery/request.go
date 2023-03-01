package delivery

import "clean-arch/features/users"

type UserRequest struct {
	Id       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

func requestToEntity(dataRequest UserRequest) users.UserEntity {
	return users.UserEntity{
		Id:       uint(dataRequest.Id),
		Name:     dataRequest.Name,
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
		Address:  dataRequest.Address,
		Role:     dataRequest.Role,
	}
}
