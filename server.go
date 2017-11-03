package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var r = gin.Default()
var gen = Generate
var add = addUser
var db *gorm.DB
var run = r.Run
var get = r.GET
var st = start

func main() {
	db = st("sqlite3", "test.db")
	get("/newUser", newUser)
	run()
}

func newUser(c *gin.Context) {
	var guid = gen(6)
	add(db, guid)
	c.JSON(200, gin.H{
		"userID": guid,
	},
	)
}
