package main

import (
	"log"

	// //iniciar o servidor e o banco de dados
	// _ "eco/src/db"
	// //Inicia a conexão com as rotas

	_ "challenge_fullstack/src/db"
	"challenge_fullstack/src/routes"
)

func main() {
	log.Print("API started")

	routes.LoadRoutes()
}
