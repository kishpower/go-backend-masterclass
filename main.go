package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kishpower/simplebank/api"
	db "github.com/kishpower/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5555/simple_bank?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func main() {
	conPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("couldn't connect with the database", err)
	}
	store := db.NewStore(conPool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("can't start the server", err)
	}
}
