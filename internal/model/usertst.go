package model

import (
	"testing"
)

func Testuser(t *testing.T) *User {
	return &User{
		ID:        "1",
		Firstname: "ivan",
		Lastname:  "Ivanov",
		Email:     "user@test.com",
		Age:       "18",
	}
}
