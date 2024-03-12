package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/physio-notes/controllers"
	middleware "github.com/guilherm5/physio-notes/middleware"
)

func HelloMiddleware(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.CheckHeader())
	api.POST("hello", controllers.Hello)
}
