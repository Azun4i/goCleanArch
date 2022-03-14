package repository

import (
	"database/sql"
	"goCleanArch/internal/model"
)

type Sqlstore struct {
	db *sql.DB
}

func NewSqlstore(db *sql.DB) Repository {
	return &Sqlstore{
		db: db,
	}
}

func (s Sqlstore) Create(u *model.User) error {
	//TODO implement me
	return nil
}

func (s Sqlstore) Delete(u *model.User) error {
	//TODO implement me
	return nil
}

func (s Sqlstore) Edit(u *model.User) error {
	//TODO implement me
	return nil
}

func (s Sqlstore) FindById(s2 string) (*model.User, error) {
	//TODO implement me
	return nil, nil
}
