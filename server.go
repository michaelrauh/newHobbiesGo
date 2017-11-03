package main

import (
	"github.com/gin-gonic/gin"
)

var r = gin.Default()
var gen = Generate

func main() {
	r.GET("/newUser", newUser)
	r.Run()
}

func newUser(c *gin.Context) {
	var guid = gen(6)
	c.JSON(200, gin.H{
		"userID": guid,
	},
	)
}
