package main

import (
	"github.com/guirialli/dolar-api/database"
	"github.com/guirialli/dolar-api/server"
)

func main() {
	database.Migrate()
	server.ServerAndListen()
}
