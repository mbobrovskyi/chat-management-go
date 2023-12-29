package user

import (
	"time"
)

type User interface {
	GetId() uint64
	GetEmail() string
	GetFirstName() string
	GetLastName() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type user struct {
	id        uint64
	email     string
	fistName  string
	lastName  string
	createdAt time.Time
	updatedAt time.Time
}

func (u *user) GetId() uint64 {
	return u.id
}

func (u *user) GetEmail() string {
	return u.email
}

func (u *user) GetFirstName() string {
	return u.fistName
}

func (u *user) GetLastName() string {
	return u.lastName
}

func (u *user) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *user) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func NewUser(
	id uint64,
	email,
	firstName,
	lastName string,
	createdAt,
	updatedAt time.Time,
) User {
	return &user{
		id:        id,
		email:     email,
		fistName:  firstName,
		lastName:  lastName,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
