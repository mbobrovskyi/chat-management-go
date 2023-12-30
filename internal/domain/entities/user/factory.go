package user

func New(id uint64, firstName string, lastName string, image string) User {
	return User{id: id, firstName: firstName, lastName: lastName, image: image}
}
