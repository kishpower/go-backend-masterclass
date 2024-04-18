package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kishpower/simplebank/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testPool *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config :", err)
	}
	testPool, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("couldn't connect with the database", err)
	}

	testQueries = New(testPool)

	os.Exit(m.Run())
}
