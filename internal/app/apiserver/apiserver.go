package apiserver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"goCleanArch/internal/repository"
	"log"
	"net/http"
)

//Start ...
func Start(config *Config) error {
	db, err := newDbConect(config.DatabaseUrl)
	if err != nil {
		log.Fatal("can't ")
	}
	defer db.Close()

	store := repository.New(db)

	s := newserver(*store)
	return http.ListenAndServe(config.BindAddr, nil)
}

//newDbConect get connect to database or err
func newDbConect(DatabaseUrl string) (db *sql.DB, err error) {

	db, err = sql.Open("postgres", DatabaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
