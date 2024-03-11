package routes

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/guilherm5/physio-notes/Middleware"
	"github.com/guilherm5/physio-notes/controllers"
)

func Cadastro(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.CheckHeader())
	api.POST("cadastro", controllers.Cadastro)
}
