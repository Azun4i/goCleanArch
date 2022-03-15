package apiserver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"goCleanArch/internal/repository"
	"log"
	"net/http"
)

//Start app
func Start(config *Config) error {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		log.Fatal("can't ", err)
	}
	defer db.Close()

	repo := repository.NewSqlstore(db)

	s := newserver(repo) //////// создаем сервер с роутерами

	log.Println(config.BindAddr)
	return http.ListenAndServe(config.BindAddr, s) /// случаем сервер
}

//newDB
// get connect to database or err
func newDB(DatabaseUrl string) (db *sql.DB, err error) {

	db, err = sql.Open("postgres", DatabaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
