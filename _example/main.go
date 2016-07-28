package main

import (
	"github.com/wantedly/_example/db"
	"github.com/wantedly/_example/server"
)

// main ...
func main() {
	database := db.Connect()
	s := server.Setup(database)
	s.Run(":8080")
}
