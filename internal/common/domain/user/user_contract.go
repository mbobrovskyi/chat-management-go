package user

import "time"

type Contract interface {
	GetCurrentUser(token string)
}

type contract struct{}

func (c *contract) GetCurrentUser(token string) User {
	return NewUser(1, "alice@gmail.com", "Alice", "Alison", time.Now(), time.Now())
}
