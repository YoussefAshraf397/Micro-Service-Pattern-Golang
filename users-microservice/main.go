package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error

func main() {

	connection()
	Migrate()

	g := gin.Default()

	posts := g.Group("users")
	{
		//posts.GET("/", getpost)
		//posts.GET("/my-posts", MyPosts)
		posts.POST("/login", LoginUser)
	}

	g.Run(":7070")
}

//func getpost(g *gin.Context) {
//	g.JSON(200, gin.H{
//		"message": "data from posts microservice",
//		"status":  "200",
//		"data":    "",
//	})
//}
func LoginUser(g *gin.Context) {
	var login Login
	if err := g.ShouldBindBodyWith(&login, binding.JSON); err != nil {
		g.JSON(400, gin.H{
			"message": "Please enter email and password",
			"status":  "false",
			"data":    "",
		})
		return
	}
	var user User
	db.Where("email =? ", login.Email).Where("password =? ", login.Password).Find(&user)
	if user.ID == 0 {
		g.JSON(400, gin.H{
			"message": "No found user by this email or password",
			"status":  "false",
			"data":    "",
		})
		return
	}

	g.JSON(200, gin.H{
		"message": "Welcome in user service",
		"status":  "true",
		"data":    user,
	})
}
