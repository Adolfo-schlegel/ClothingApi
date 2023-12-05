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
	IdClient primitive.ObjectID `json:"idClient" bson:"idClient"`
	Clothes  []Clothe           `json:"clothes" bson:"clothes"`
}

type Clothe struct {
	ID                    primitive.ObjectID `json:"_id" bson:"_id"`
	Date                  string             `json:"date" bson:"date"`
	Count                 string             `json:"count" bson:"count"`
	EmbroideryInstruction string             `json:"embroideryInstruction" bson:"embroideryInstruction"`
	Descripcion           string             `json:"descripcion" bson:"descripcion"`
	Cost                  int                `json:"cost" bson:"cost"`
	Color                 string             `json:"color" bson:"color"`
}

type IntegralClient struct {
	ID       primitive.ObjectID `json:"id" bson:"id"`
	IdClient string             `json:"idCliente" bson:"idClient"`
}
