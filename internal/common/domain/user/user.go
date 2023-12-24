package user

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/entity"
	"time"
)

type User interface {
	entity.Entity[User]

	GetEmail() string
	GetFirstName() string
	GetLastName() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type user struct {
	entity.Entity[User]

	Email     string
	FistName  string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(
	id uint64,
	email,
	firstName,
	lastName string,
	createdAt,
	updatedAt time.Time,
) *user {
	return &user{
		Entity:    entity.New[User](id),
		Email:     email,
		FistName:  firstName,
		LastName:  lastName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
