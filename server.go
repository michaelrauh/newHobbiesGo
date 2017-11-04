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
var addH = addHobby
var allH = allHobbies

func main() {
	db = st("sqlite3", "test.db")
	addH(db, "text")
	get("/newUser", newUser)
	get("/hobbies", hobbies)

	run()
}

func hobbies(c *gin.Context) {
	c.JSON(200, gin.H{
		"hobbies": allH(db),
	},
	)
}

func newUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"userID": addUniq(db),
	},
	)
}
