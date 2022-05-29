package main

import (
	"github.com/gin-gonic/gin"
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
		serviceUrl := segment[0]

		g.JSON(200, gin.H{
			"message": "We  found the service",
			"status":  "True",
			"data":    serviceUrl,
		})
	}
}

func ServicesList() map[string]string {
	m := make(map[string]string)
	m["users"] = "http://localhost:7070"
	m["posts"] = "http://localhost:6060"

	return m
}
