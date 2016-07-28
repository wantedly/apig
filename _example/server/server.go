package server

import (
	"github.com/wantedly/apig/_example/middleware"
	"github.com/wantedly/apig/_example/router"

  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Setup(db *gorm.DB) *gin.Engine {
  r := gin.Default()
  r.Use(middleware.SetDBtoContext(db))
  router.Initialize(r)
  return r
}
