package main

import (
	"github.com/gin-gonic/gin"
	)

func newUser(c *gin.Context) {
  c.JSON(200, gin.H{
			"userID":   "abc123",
		},
	)
}
