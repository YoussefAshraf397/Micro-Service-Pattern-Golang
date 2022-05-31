package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json: "name"`
	Email    string `json: "email"`
	Token    string `json: "token"`
	Password string `json: "password"`
}

type Login struct {
	Email    string `json: "email" binding:"required`
	Password string `json: "password" binding:"required`
}
