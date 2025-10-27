package main

import (
	"database/sql"
	"log"

	"github.com/junasliang/go_simplebank/api"
	db "github.com/junasliang/go_simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	DBSource     = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	severAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(severAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
