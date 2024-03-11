package controllers

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.JSON(200, "Hello fisio")
}
