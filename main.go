package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/physio-notes/routes"
)

func main() {
	router := gin.Default()
	routes.Cadastro(router)
	routes.Login(router)
	routes.HelloMiddleware(router)
	router.Run(":8080")
}
