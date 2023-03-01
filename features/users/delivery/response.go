package delivery

import "clean-arch/features/users"

type UserResponse struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Role    string `json:"role"`
}

func entityToResponse(userEntity users.UserEntity) UserResponse {
	return UserResponse{
		Id:      userEntity.Id,
		Name:    userEntity.Name,
		Email:   userEntity.Email,
		Address: userEntity.Address,
		Role:    userEntity.Role,
	}
}

func listEntityToResponse(userEntities []users.UserEntity) []UserResponse {
	var userResponse []UserResponse
	for _, v := range userEntities {
		userResponse = append(userResponse, entityToResponse(v))
	}
	return userResponse
}
