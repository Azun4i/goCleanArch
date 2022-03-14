package repository

import "goCleanArch/internal/model"

type Repository interface {
	Create(u *model.User) error
	Delete(u *model.User) error
	Edit(u *model.User) error
	FindById(str string) (*model.User, error)
}
