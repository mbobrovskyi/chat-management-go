package user

import "time"

type Contract interface {
	GetCurrentUser(token string) (User, error)
}

type contract struct{}

func (c *contract) GetCurrentUser(token string) (User, error) {
	// TODO: Get user from management service
	return NewUser(1, "alice@gmail.com", "Alice", "Alison", time.Now(), time.Now()), nil
}

func NewContract() Contract {
	return &contract{}
}
