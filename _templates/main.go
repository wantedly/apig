package main

import (
	"github.com/USERNAME/APISERVER/db"
	"github.com/USERNAME/APISERVER/server"
)

// main ...
func main() {
	database := db.Connect()
	s := server.Setup(database)
	s.Run(":8080")
}
