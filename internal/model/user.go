package model

type User struct {
	ID        string
	Firstname string
	Lastname  string
	Email     string
	Age       string
}

func NewUser() *User {
	return &User{}
}
