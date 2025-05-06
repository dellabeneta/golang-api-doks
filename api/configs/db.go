package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// Configuração do cliente MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	
	// Contexto com timeout para a conexão
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Conecta ao MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verifica a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Conectado ao MongoDB com sucesso!")
	return client
}

// Instância do cliente MongoDB
var DB *mongo.Client = ConnectDB()

// GetCollection retorna uma coleção do MongoDB
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("peopledb").Collection(collectionName)
	return collection
}
