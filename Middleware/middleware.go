package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CheckHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.GetHeader("Content-Type")
		if contentType != "application/json" {
			log.Println("Content-Type incorreto")
			c.JSON(400, gin.H{
				"error": "Erro no header da requisição, procure o desenvolvedor",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
