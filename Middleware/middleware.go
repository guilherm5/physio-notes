package middleware

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := os.Getenv("secret")

		contentType := c.GetHeader("Content-Type")
		if contentType != "application/json" {
			log.Println("Content-Type incorreto")
			c.JSON(400, gin.H{
				"error": "Erro no header da requisição, procure o desenvolvedor",
			})
			c.AbortWithStatus(401)
			return
		}

		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			log.Println("Cliente mandando header authorization vazio")
			c.JSON(401, gin.H{
				"Erro: ": "Erro ao preencher sua requisição, entre em contato com o desenvolvedor",
			})
			c.AbortWithStatus(401)
			return
		}

		token, err := jwt.Parse(authorization, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil

		})
		if err != nil || !token.Valid {
			log.Println("Token inválido", err)
			c.JSON(401, gin.H{
				"Erro: ": "Sua autenticação parece estar inválida, por favor, realize novamente o login",
			})
			c.AbortWithStatus(401)
			return
		}
	}
}
