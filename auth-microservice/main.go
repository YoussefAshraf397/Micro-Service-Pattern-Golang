package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var authorizationToken string = "authtoken"

var db *gorm.DB = nil
var err error

func main() {

	connection()
	Migrate()

	g := gin.Default()

	g.GET("/", Welcome)
	g.POST("removeStoreToken", RemoveStoreToken)
	g.POST("logout", Logout)

	g.Use(Services())

	g.Run(":5050")
}

func Welcome(g *gin.Context) {
	g.JSON(200, gin.H{
		"message": "done",
		"status":  "200",
		"data":    "",
	})
	return
}

func RemoveStoreToken(g *gin.Context) {
	authorization := g.GetHeader("auth_token")
	fmt.Println("auth token in authnicroservice: ", authorization)

	var removeToken RemoveToken
	if err := g.ShouldBindJSON(&removeToken); err != nil {
		fmt.Println("is error1")
		g.JSON(400, gin.H{
			"message": "Token is not valid",
			"status":  "false",
			"data":    "",
		})

	}
	fmt.Println("no error1")

	if removeToken.AuthToken != authorizationToken {
		fmt.Println("is error2")

		g.JSON(400, gin.H{
			"message": "AuthToken is not valid",
			"status":  "false",
			"data":    "",
		})
	}
	fmt.Println("remove token: ", removeToken.UserId)
	db.Where("user_id =? ", removeToken.UserId).Unscoped().Delete(&Token{}) //Hard Delete
	token := Token{
		Token:  removeToken.UserToken,
		UserId: removeToken.UserId,
	}
	db.Create(&token)
	g.JSON(200, gin.H{
		"message": "Done token is updated",
		"status":  "true",
		"data":    "",
	})

}

func Logout(g *gin.Context) {
	token := g.GetHeader("authorization")
	fmt.Println("logout token is: ", token)
	if len(token) == 0 {
		g.JSON(401, gin.H{
			"message": "Yous should put the token authorization in headers.",
			"status":  "false",
			"data":    "",
		})
		return
		g.Abort()
	}
	var findToken Token
	db.Where("token =? ", token).First(&findToken)
	if findToken.ID == 0 {
		g.JSON(401, gin.H{
			"message": "Not Authorize token dost match any tokens ",
			"status":  "false",
			"data":    "",
		})
		return
		g.Abort()
	}

	db.Where("token =? ", token).Unscoped().Delete(&Token{}) //Hard Delete

	g.JSON(200, gin.H{
		"message": "Done Logout",
		"status":  "true",
		"data":    "",
	})

}
