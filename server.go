package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var r = gin.Default()

var db *gorm.DB
var run = r.Run
var get = r.GET
var st = start
var addUniq = addUnique

func main() {
	db = st("sqlite3", "test.db")
	get("/newUser", newUser)
	run()
}

func newUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"userID": addUniq(db),
	},
	)
}
