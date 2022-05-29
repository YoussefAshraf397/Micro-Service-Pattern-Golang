package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Services() gin.HandlerFunc {
	return func(g *gin.Context) {
		url := g.Request.RequestURI
		segmenats := strings.Split(url, "/")

		services := ServicesList()
		if ok := services[segmenats[1]]; !ok {
			g.JSON(404, gin.H{
				"message": "We not found the service",
				"status":  "false",
				"data":    "",
			})
			return
			g.Abort()
		}
		g.JSON(200, gin.H{
			"message": "We  found the service",
			"status":  "True",
			"data":    "",
		})
	}
}

func ServicesList() map[string]bool {
	m := make(map[string]bool)
	m["users"] = true
	m["posts"] = true

	return m
}
