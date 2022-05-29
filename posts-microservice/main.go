package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error

func main() {

	g := gin.Default()

	posts := g.Group("posts")
	{
		posts.GET("/", getpost)
	}

	g.Run(":6060")
}

func getpost(g *gin.Context) {
	g.JSON(200, gin.H{
		"message": "data from posts microservice",
		"status":  "200",
		"data":    "",
	})
}
