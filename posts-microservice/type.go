package main

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `json: "title"`
	Des   string `json: "des"`

	UserId int `json: "userId"`
}
