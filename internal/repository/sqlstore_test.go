package repository_test

import (
	"os"
	"testing"
)

var databaseURL string

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5432 user=azunai password=0000 dbname=apicl_test sslmode=disable"
	}
	os.Exit(m.Run()) // run возврадает код ошибки
}
