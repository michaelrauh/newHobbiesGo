package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/newUser", newUser)
	r.Run()
}

var gen = Generate

func newUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"userID": gen(6),
	},
	)
}
