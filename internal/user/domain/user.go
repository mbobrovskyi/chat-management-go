package domain

type User interface {
	GetId() uint64
	GetFistName() string
	GetLastName() string
	GetImage() string
}

type user struct {
	id        uint64
	firstName string
	lastName  string
	image     string
}

func (u *user) GetId() uint64 {
	return u.id
}

func (u *user) GetFistName() string {
	return u.firstName
}

func (u *user) GetLastName() string {
	return u.lastName
}

func (u *user) GetImage() string {
	return u.image
}

func NewUser(id uint64, firstName string, lastName string, image string) User {
	return &user{id: id, firstName: firstName, lastName: lastName, image: image}
}
