package data_mappers

import (
	"github.com/mbobrovskyi/chat-management-go/internal/common/data_contracts"
	"github.com/mbobrovskyi/chat-management-go/internal/user/domain"
)

func UserFromDTO(input data_contracts.UserDTO) domain.User {
	return domain.NewUser(input.Id, input.FirstName, input.LastName, input.Image)
}

func UserToDTO(input domain.User) data_contracts.UserDTO {
	return data_contracts.UserDTO{
		Id:        input.GetId(),
		FirstName: input.GetFistName(),
		LastName:  input.GetLastName(),
		Image:     input.GetImage(),
	}
}
