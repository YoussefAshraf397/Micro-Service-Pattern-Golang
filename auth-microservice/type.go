package main

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	Token  string `json: "token"`
	UserId int    `json: "userId"`
}

type RemoveToken struct {
	AuthToken string `json: "auth_token" binding:"required`
	UserToken string `json: "user_token" binding:"required`
	UserId    int    `json: "user_id" binding:"required`
}
