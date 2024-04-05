package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5555/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("couldn't connect with the database")
	}

	testQueries = New(dbConn)

	os.Exit(m.Run())
}
