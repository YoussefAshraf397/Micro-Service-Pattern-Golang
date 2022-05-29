package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
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

		//log.Fatalln(NewUrl)

		Get(g, NewUrl)
		fmt.Println("url is : ", NewUrl)

		//g.JSON(200, gin.H{
		//	"message": "We  found the service",
		//	"status":  "True",
		//	"data":    serviceUrl,
		//})
	}
}

func Get(g *gin.Context, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error in your request")
	}

	for key, _ := range g.Request.Header {
		req.Header.Set(key, g.Request.Header.Get(key))
	}

	client := &http.Client{Timeout: time.Second * 10}

	response, err := client.Do(req)

	if err != nil {
		fmt.Println("error in do request")
	}
	defer response.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)
	if len(result) == 0 {
		g.JSON(500, gin.H{
			"message": "Not found any data in response",
			"status":  "false",
			"data":    "",
		})
		return
	}
	g.JSON(response.StatusCode, result)
	return

}

func ServicesList() map[string]string {
	m := make(map[string]string)
	m["users"] = "http://127.0.0.1:7070"
	m["posts"] = "http://127.0.0.1:6060"

	return m

}
