package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func Services() gin.HandlerFunc {
	return func(g *gin.Context) {
		url := g.Request.RequestURI
		segmenats := strings.Split(url, "/")
		segment := strings.Split(segmenats[1], "?")

		services := ServicesList()
		if ok := services[segment[0]]; ok == "" {
			g.JSON(404, gin.H{
				"message": "We not found the service",
				"status":  "false",
				"data":    "",
			})
			return
			g.Abort()
		}
		serviceUrl := services[segment[0]]
		NewUrl := serviceUrl + url

		var token Token

		authList := AuthRouteList()
		for route, _ := range authList {
			if strings.Contains(url, route) {
				authorization := g.GetHeader("authorization")
				if authorization == "" {
					g.JSON(401, gin.H{
						"message": "Not Authorize No token in header",
						"status":  "false",
						"data":    "",
					})
					return
					g.Abort()
				}

				db.Where("token =? ", authorization).First(&token)
				if token.ID == 0 {
					g.JSON(401, gin.H{
						"message": "Not Authorize token dost match any tokens ",
						"status":  "false",
						"data":    "",
					})
					return
					g.Abort()
				}

				// Authorized User
				g.Request.Header.Set("USER_ID", strconv.Itoa(token.UserId))

			}
		}

		//log.Fatalln(NewUrl)

		methode := strings.ToLower(g.Request.Method)
		switch methode {
		case "post":
			Post(g, NewUrl)
			break
		case "get":
			Get(g, NewUrl)

		}

		//Get(g, NewUrl)
		//fmt.Println("url is : ", NewUrl)

		//g.JSON(200, gin.H{
		//	"message": "We  found the service",
		//	"status":  "True",
		//	"data":    serviceUrl,
		//})
	}
}

func ServicesList() map[string]string {
	m := make(map[string]string)
	m["users"] = "http://127.0.0.1:9090"
	m["posts"] = "http://127.0.0.1:6060"

	return m
}

func AuthRouteList() map[string]bool {
	m := make(map[string]bool)
	m["my-posts"] = true

	return m
}
