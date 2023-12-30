package domain

type Contract interface {
	GetCurrent(token string) (User, error)
	GetAll() ([]User, error)
}
