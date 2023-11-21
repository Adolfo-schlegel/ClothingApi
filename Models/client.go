package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID       primitive.ObjectID `json:"id" bson:"id"`
	Name     string             `json:"name"`
	Contact  string             `json:"contact"`
	Date     string             `json:"date"`
	Discount string             `json:"discount"`
	State    bool               `json:"state"`
	Image    string             `json:"image"`
	Rating   int                `json:"rating"`
	Type     string             `json:"type"`
}

type ParcialClient struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	IdClient string             `json:"idClient" bson:"idClient"`
	Clothes  []Clothe           `json:"clothes" bson:"clothes"`
}

type Clothe struct {
	Date                  string `json:"date" bson:"date"`
	Count                 string `json:"count"`
	EmbroideryInstruction string `json:"embroideryInstruction"`
	Descripcion           string `json:"descripcion"`
	Cost                  int    `json:"cost"`
	Color                 string `json:"color"`
}

type IntegralClient struct {
	ID       primitive.ObjectID `json:"id" bson:"id"`
	IdClient string             `json:"idCliente" bson:"idClient"`
}
