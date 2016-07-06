package main

import (
	"github.com/wantedly/api-server-generator/examples/simple/db"
	"github.com/wantedly/api-server-generator/examples/simple/server"
)

// main ...
func main() {
	database := db.Connect()
	s := server.Setup(database)
	s.Run(":8080")
}
