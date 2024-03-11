package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/physio-notes/database"
	"github.com/guilherm5/physio-notes/models"
	"golang.org/x/crypto/bcrypt"
)

var DB = database.DatabasePhysio()

func Cadastro(c *gin.Context) {
	data := models.CadastroFisio{}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("Erro ao ler json enviado na requisição", err)
		c.JSON(400, gin.H{
			"Erro: ": "Parece que existe um erro no formato da requisição, entre em contato com o desenvolvedor",
		})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data.Senha), 10)
	if err != nil {
		log.Println("Erro ao gerar hash da senha do fisioterapeura", err)
		c.JSON(400, gin.H{
			"Erro: ": "Parece que existe um erro interno ao processar as informações de cadastro, por favor, entre em contato com o desenvolvedor",
		})
		return
	}
	cmd, err := DB.Prepare("INSERT INTO fisioterapeuta (nome_fisioterapeuta, email_fisioterapeuta, senha_fisioterapeuta) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println("Erro ao preparar a instrução (insert) SQL:", err)
		c.JSON(400, gin.H{
			"Erro: ": "Aconteceu um erro ao preparar o cadastro de usuário, por favor, entre em contato com o desenvolvedor",
		})
		return
	}
	defer cmd.Close()

	_, err = cmd.Exec(data.Nome, data.Email, string(password))
	if err != nil {
		log.Println("Erro ao executar a instrução (insert) SQL:", err)
		c.JSON(400, gin.H{
			"Erro: ": "Aconteceu um erro ao inserir seu cadastro, por favor, entre em contato com o desenvolvedor",
		})
		return
	}
	c.JSON(200, gin.H{
		"Mensagem:": "Sucesso ao realizar cadastro",
	})
}
