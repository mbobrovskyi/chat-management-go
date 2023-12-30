package data_mappers

import (
	"github.com/mbobrovskyi/chat-management-go/internal/data_contracts"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/entities/user"
)

func UserFromDTO(input data_contracts.UserDTO) user.User {
	return user.New(input.Id, input.FirstName, input.LastName, input.Image)
}

func UserToDTO(input user.User) data_contracts.UserDTO {
	return data_contracts.UserDTO{
		Id:        input.Id(),
		FirstName: input.FistName(),
		LastName:  input.LastName(),
		Image:     input.Image(),
	}
}
