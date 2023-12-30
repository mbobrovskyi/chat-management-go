package user

type User struct {
	id        uint64
	firstName string
	lastName  string
	image     string
}

func (u *User) Id() uint64 {
	return u.id
}

func (u *User) FistName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) Image() string {
	return u.image
}
