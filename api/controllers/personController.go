package controllers

import (
	"context"
	"encoding/json"
	"golang-api/configs"
	"golang-api/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var personCollection *mongo.Collection = configs.GetCollection(configs.DB, "pessoas")

// CreatePerson cria um novo cadastro de pessoa
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var person models.Person
	
	// Decodifica o corpo da requisição
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"status": "error", "message": "Dados inválidos"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Verifica se o CPF já existe
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	filter := bson.M{"cpf": person.CPF}
	count, err := personCollection.CountDocuments(ctx, filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"status": "error", "message": "Erro ao verificar CPF"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	if count > 0 {
		w.WriteHeader(http.StatusConflict)
		response := map[string]string{"status": "error", "message": "CPF já cadastrado"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Insere a nova pessoa
	result, err := personCollection.InsertOne(ctx, person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"status": "error", "message": "Erro ao cadastrar pessoa"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"status": "success", 
		"message": "Pessoa cadastrada com sucesso",
		"id": result.InsertedID,
	}
	json.NewEncoder(w).Encode(response)
}

// GetPersonByCPF busca pessoa pelo CPF
func GetPersonByCPF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Obtém o CPF da URL
	params := mux.Vars(r)
	cpf := params["cpf"]
	
	var person models.Person
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Busca a pessoa pelo CPF
	filter := bson.M{"cpf": cpf}
	err := personCollection.FindOne(ctx, filter).Decode(&person)
	
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			response := map[string]string{"status": "error", "message": "Pessoa não encontrada"}
			json.NewEncoder(w).Encode(response)
			return
		}
		
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"status": "error", "message": "Erro ao buscar pessoa"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}