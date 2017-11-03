package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type user struct {
	gorm.Model
	GUID string
}

var db *gorm.DB
var err error
var r = gin.Default()

func main() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&user{})

	r.GET("/newUser", newUser)
	r.Run()
}

var gen = Generate

func newUser(c *gin.Context) {
	var guid = gen(6)
	db.Create(&user{GUID: guid})
	c.JSON(200, gin.H{
		"userID": guid,
	},
	)
}
