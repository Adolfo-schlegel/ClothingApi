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
	Boxes    []Box              `json:"boxes" bson:"clothes"`
}

type Box struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Date             string             `json:"date" bson:"date"`
	Count            string             `json:"count" bson:"count"`
	EmbroideryImages []Image            `json:"embroideryImages" bson:"embroideryImages"`
	Descripcion      string             `json:"descripcion" bson:"descripcion"`
	Cost             int                `json:"cost" bson:"cost"`
	Color            string             `json:"color" bson:"color"`
}
type Image struct {
	Name    string `json:"name" bson:"name"`
	Content string `json:"content" bson:"content"`
}

type IntegralClient struct {
	ID       primitive.ObjectID `json:"id" bson:"id"`
	IdClient string             `json:"idCliente" bson:"idClient"`
}

type Order struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	IdClient   primitive.ObjectID `json:"idClient" bson:"idClient"`
	NameClothe string             `json:"nameClothe" bson:"nameClothe"`
	Waist      float32            `json:"waist" bson:"waist"`
	Color      string             `json:"color" bson:"color"`
	Embroidery string             `json:"embroidery" bson:"embroidery"`
	CostPrice  float32            `json:"costPrice" bson:"costPrice"`
	SalePrice  float32            `json:"salePrice" bson:"salePrice"`
}

type Clothe struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"` // TODO
}
