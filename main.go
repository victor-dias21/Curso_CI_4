package main

import (
	"fmt"
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/routes"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Carregar variáveis do .env primeiro
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	// 2. Pegar variáveis de ambiente
	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	dbname := os.Getenv("DBNAME")
	dbport := os.Getenv("DBPORT")
	port := os.Getenv("PORT")

	// Apenas para debug (não exiba senha em produção!)
	fmt.Println("Conectando em:", host, "com usuário:", user)
	fmt.Println("Banco:", dbname, "Porta DB:", dbport)
	fmt.Println("Aplicação rodando na porta:", port)

	// 3. Conectar ao banco usando as envs
	database.ConectaComBancoDeDados()

	// 4. Iniciar rotas
	routes.HandleRequest()
}
