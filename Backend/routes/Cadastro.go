package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/physio-notes/controllers"
)

func Cadastro(c *gin.Engine) {
	api := c.Group("api")
	api.POST("cadastro", controllers.Cadastro)
}
