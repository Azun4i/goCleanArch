package apiserver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"goCleanArch/internal/repository"
	"goCleanArch/internal/usecases"
	"log"
	"net/http"
)

//Start app
func Start(config *Config) error {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		log.Fatal("can't connect to db ", err)
	}
	defer db.Close()

	// получаем репозитории
	repo := repository.NewSqlstore(db)

	// получаем в use case repository
	us := usecases.NewUseCase(repo)

	// сервис получает usecase
	s := Newserver(us)

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
