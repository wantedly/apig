package db

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Connect() *gorm.DB {
	dir := filepath.Dir("db/database.db")
	db, err := gorm.Open("sqlite3", dir+"/database.db")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	return db
}

func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}
