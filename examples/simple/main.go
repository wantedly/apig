package main

import (
	"github.com/wantedly/api-server-generator/examples/simple/db"
	"github.com/wantedly/api-server-generator/examples/simple/models"
	"github.com/wantedly/api-server-generator/examples/simple/server"
)

// main ...
func main() {
	database := db.Connect()

	// TODO(awakia): Consider position of automigrate
	database.AutoMigrate(&models.User{})

	s := server.Setup(database)
	s.Run(":8080")
}
