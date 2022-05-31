package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/HttpRequest"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error

var authorizationToken string = "authtoken"
var authUrl string = "http://localhost:5050"

func main() {

	connection()
	Migrate()

	g := gin.Default()

	posts := g.Group("users")
	{
		posts.POST("/login", LoginUser)
		//posts.POST("/logout", LogoutUser)

	}

	g.Run(":9090")
}

func LoginUser(g *gin.Context) {
	var login Login
	if err := g.ShouldBindJSON(&login); err != nil {
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

	token, _ := bcrypt.GenerateFromPassword([]byte(user.Email+"*()*654654654"), 3)
	user.Token = string(token)
	db.Save(&user)

	fmt.Println("user id in usermicroservice: ", user.ID)
	req := HttpRequest.NewRequest()
	req.JSON().Post(authUrl+"/removeStoreToken", map[string]interface{}{
		"AuthToken": authorizationToken,
		"UserToken": user.Token,
		"UserId":    user.ID,
	})

	g.JSON(200, gin.H{
		"message": "Welcome in user service",
		"status":  "true",
		"data":    user,
	})
}

//func LogoutUser(g *gin.Context) {
//
//}
