package data_contracts

type UserDTO struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Image     string `json:"image"`
}
