package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DatabasePhysio() *sql.DB {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Erro ao carregar variaveis de embiente na funcao de banco de dados", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", os.Getenv("strBD"))
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados", err)

	}
	log.Println("Sucesso ao logar no banco de dados")

	return db
}
