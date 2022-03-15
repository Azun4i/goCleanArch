package repository

import (
	"database/sql"
	"fmt"
	"goCleanArch/internal/model"
	"log"
)

type Sqlstore struct {
	db *sql.DB
}

func NewSqlstore(db *sql.DB) Repository {
	return &Sqlstore{
		db: db,
	}
}

//Create
//inset model in to database or get err
func (s Sqlstore) Create(u *model.User) error {
	if _, err := s.db.Query("INSERT INTO authors (uuid, firstname, lastname,email, age) VALUES ($1,$2,$3,$4,$5)",
		u.ID, u.Firstname, u.Lastname, u.Email, u.Age,
	); err != nil {
		return err
	}
	return nil
}

func (s Sqlstore) Delete(id string) error {
	del, err := s.db.Prepare("DELETE FROM authors WHERE uuid=$1")
	if err != nil {

		return err
	}
	del.Exec(id)
	return nil
}

func (s Sqlstore) Edit(u *model.User) error {
	log.Println(u.ID, u.Firstname, u.Email, u.Age)
	if _, err := s.db.Exec("UPDATE authors SET firstname=$1, lastname=$2,email=$3, age=$4 WHERE uuid=$5",
		u.Firstname, u.Lastname, u, u.Email, u.Age,
		u.ID,
	); err != nil {
		return err
	}
	return nil
}

func (s Sqlstore) FindById(id string) (*model.User, error) {
	u := &model.User{}
	fmt.Println(id)
	if err := s.db.QueryRow("SELECT uuid, firstname, lastname, email, age FROM authors WHERE uuid=$1", id).Scan(
		&u.ID,
		&u.Firstname,
		&u.Lastname,
		&u.Email,
		&u.Age,
	); err != nil {
		return nil, err
	}
	return u, nil
}
