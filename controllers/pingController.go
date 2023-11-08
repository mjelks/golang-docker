package controllers

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {

	// return it
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}
