package main

import (
	"database/sql"
	"log"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/junasliang/go_simplebank/api"
	db "github.com/junasliang/go_simplebank/db/sqlc"
	"github.com/junasliang/go_simplebank/db/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
