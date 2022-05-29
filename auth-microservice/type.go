package main

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	Token  string `json: "token"`
	UserId int    `json: "userId"`
}
