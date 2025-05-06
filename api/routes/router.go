package routes

import (
	"golang-api/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	// Rota para criar pessoa
	router.HandleFunc("/api/pessoas", controllers.CreatePerson).Methods("POST")
	
	// Rota para buscar pessoa por CPF
	router.HandleFunc("/api/pessoas/{cpf}", controllers.GetPersonByCPF).Methods("GET")
}
