package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/physio-notes/controllers"
)

func Login(c *gin.Engine) {
	api := c.Group("api")
	api.POST("login", controllers.Login)
}
