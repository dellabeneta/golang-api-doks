package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Person define o modelo de pessoa para o cadastro
type Person struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nome      string             `json:"nome" bson:"nome" validate:"required"`
	Sobrenome string             `json:"sobrenome" bson:"sobrenome" validate:"required"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	CPF       string             `json:"cpf" bson:"cpf" validate:"required"`
}