package model

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Email     string
	Age       uint
}

func NewUser() *User {
	return &User{}
}
