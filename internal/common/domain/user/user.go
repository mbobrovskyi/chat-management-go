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

	Email     string    `json:"email"`
	FistName  string    `json:"fistName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *user) GetEmail() string {
	return u.Email
}

func (u *user) GetFirstName() string {
	return u.FistName
}

func (u *user) GetLastName() string {
	return u.LastName
}

func (u *user) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *user) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func New(
	id uint64,
	email,
	firstName,
	lastName string,
	createdAt,
	updatedAt time.Time,
) User {
	return &user{
		Entity:    entity.New[User](id),
		Email:     email,
		FistName:  firstName,
		LastName:  lastName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
