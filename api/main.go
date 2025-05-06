package main

import (
	"fmt"
	// "golang-api/configs"
	"golang-api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializa o router
	router := mux.NewRouter()
	
	// Conecta ao MongoDB
	// configs.ConnectDB()
	
	// Configura as rotas
	routes.SetupRoutes(router)
	
	// Inicia o servidor
	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}