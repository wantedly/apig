package main

import (
	"github.com/wantedly/apig/_examples/simple/db"
	"github.com/wantedly/apig/_examples/simple/models"
	"github.com/wantedly/apig/_examples/simple/server"
)

// main ...
func main() {
	database := db.Connect()

	// TODO(awakia): Consider position of automigrate
	database.AutoMigrate(&models.User{})

	s := server.Setup(database)
	s.Run(":8080")
}
