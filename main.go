package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kishpower/simplebank/api"
	db "github.com/kishpower/simplebank/db/sqlc"
	utils "github.com/kishpower/simplebank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal()
	}
	conPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("couldn't connect with the database", err)
	}
	store := db.NewStore(conPool)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can't start the server", err)
	}
}
