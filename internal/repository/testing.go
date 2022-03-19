package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
	"testing"
)

//TestDb
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper() // указывает что это тестовый медод

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}
	return db, func(s ...string) {
		if len(s) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(s, ", ")))
		}
		db.Close()
	}

}
