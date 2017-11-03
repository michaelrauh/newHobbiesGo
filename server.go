package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var r = gin.Default()
var gen = Generate
var add = addUser
var db *gorm.DB

func main() {
	db = start()
	r.GET("/newUser", newUser)
	r.Run()
}

func newUser(c *gin.Context) {
	var guid = gen(6)
	add(db, guid)
	c.JSON(200, gin.H{
		"userID": guid,
	},
	)
}
