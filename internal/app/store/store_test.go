package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == ""{
		databaseURL = "sqlserver://MyPC/?database=UriDb"
	}

	os.Exit(m.Run())
}