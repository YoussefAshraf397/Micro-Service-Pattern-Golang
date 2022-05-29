package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error

func main() {
	connection()
	db.AutoMigrate(Token{})
	g := gin.Default()

	g.GET("/", Welcome)
	g.Run(":5050")
}

func Welcome(g *gin.Context) {
	g.JSON(200, gin.H{
		"message": "done",
		"status":  "200",
		"data":    "",
	})
}
