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
