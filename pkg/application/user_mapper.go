package application

import (
	user2 "github.com/mbobrovskyi/chat-management-go/internal/contracts"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/entities/user"
)

func UserFromDTO(input user2.UserDTO) user.User {
	return user.New(input.Id, input.FirstName, input.LastName, input.Image)
}

func UserToDTO(input user.User) user2.UserDTO {
	return user2.UserDTO{
		Id:        input.Id(),
		FirstName: input.FistName(),
		LastName:  input.LastName(),
		Image:     input.Image(),
	}
}
