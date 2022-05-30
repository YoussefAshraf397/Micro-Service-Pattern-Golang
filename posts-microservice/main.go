package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error

func main() {

	connection()
	Migrate()

	g := gin.Default()

	posts := g.Group("posts")
	{
		posts.GET("/", getpost)
		posts.GET("/my-posts", MyPosts)
		posts.POST("/", getpost)
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
func MyPosts(g *gin.Context) {
	fmt.Println(g.Request.Header)
	var posts []Post
	userID := g.GetHeader("USER_ID")

	db.Where("user_id =? ", userID).Find(&posts)

	g.JSON(200, gin.H{
		"message": "You should see all your posts",
		"status":  "200",
		"data":    posts,
	})
}
