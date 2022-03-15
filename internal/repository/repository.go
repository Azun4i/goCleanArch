package repository

import "goCleanArch/internal/model"

type Repository interface {
	Create(u *model.User) error
	Delete(id string) error
	Edit(u *model.User) error
	FindById(id string) (*model.User, error)
}
