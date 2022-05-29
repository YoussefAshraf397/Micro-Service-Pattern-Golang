package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error

func main() {

	g := gin.Default()
	g.GET("/", Welcome)
	g.Run(":6060")
}

func Welcome(g *gin.Context) {
	g.JSON(200, gin.H{
		"message": "done",
		"status":  "200",
		"data":    "",
	})
}
