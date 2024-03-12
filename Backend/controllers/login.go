package controllers

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/physio-notes/models"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Erro ao carregar variaveis de embiente para login", err)
		c.Status(400)
		return
	}
	secret := os.Getenv("secret")
	credenciais := models.User{}
	login := models.User{}

	err = c.ShouldBindJSON(&credenciais)
	if err != nil {
		log.Println("Erro ao ler credenciais", err)
		c.JSON(400, gin.H{
			"Erro: ": "O sistema não consegue ler suas credenciais, entre em contato com o desenvolvedor",
		})
		return
	}
	query, err := DB.Query("SELECT * FROM fisioterapeuta WHERE email_fisioterapeuta = $1", &credenciais.Email)
	if err != nil {
		log.Println("Erro ao buscar usuário no banco de dados", err)
		c.JSON(400, gin.H{
			"Erro: ": "O sistema não consegue encontrar seu usuário, você já tem um login definido ?",
		})
		return
	}

	for query.Next() {
		err = query.Scan(&login.IDFisioterapeuta, &login.Nome, &login.Email, &login.Senha)
		if err != nil {
			log.Println("Erro ao buscar usuario para realizar login", err)
			c.JSON(400, gin.H{
				"Erro: ": "O sistema não consegue retornar os dados do seu usuário, entre em contato com o desenvolvedor",
			})
			return
		}
		c.Next()
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.Senha), []byte(credenciais.Senha))
	if err != nil {
		log.Println("Senha inválida", err)
		c.JSON(401, gin.H{
			"Erro": "Senha inválida",
		})
		return
	}
	claims := jwt.MapClaims{
		"ID":   login.IDFisioterapeuta,
		"Nome": login.Nome,
		//"Perm": login.TipoUsuario,
		"exp": time.Hour,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Erro ao gerar JWT", err)
		c.JSON(500, gin.H{
			"Erro: ": "Aconteceu um erro interno ao liberar seu acesso, entre em contato com o desenvolvedor",
		})
		return
	}

	c.JSON(200, tokenString)
}
