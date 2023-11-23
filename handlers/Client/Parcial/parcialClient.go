package ParcialClient

import (
	Model "example/src/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection = nil

func GetPartials(c *gin.Context) {
	//Find all
	cursor, err := Collection.Find(c, bson.M{})

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"\nmessage": "Error parsing JSON", "code": http.StatusInternalServerError})
		return
	}

	defer cursor.Close(c)

	var results []Model.ParcialClient

	for cursor.Next(c) {
		var result Model.ParcialClient
		err := cursor.Decode(&result)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"\nmessage": "Error parsing JSON: " + err.Error(), "code": http.StatusInternalServerError})
			return
		}
		results = append(results, result)
	}

	c.IndentedJSON(http.StatusOK, results)
}
func GetPartialById(c *gin.Context) {
	client, err := findPartial(c.Param("id"), c)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"\nmessage": "Client not found " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, client)
}

func findPartial(clientId string, c *gin.Context) (*Model.ParcialClient, error) {
	var result Model.ParcialClient

	// Create a primitive.ObjectID from a hexadecimal string
	id, err := primitive.ObjectIDFromHex(clientId)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{"idClient", id}}

	err = Collection.FindOne(c.Request.Context(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateParcial(c *gin.Context) {

	//get the client from the json request
	var client Model.ParcialClient

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Set a new ObjectID for the client

	client.ID = primitive.NewObjectID()

	//Insert the client into the database
	_, err := Collection.InsertOne(c, client)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error: " + err.Error()})
		return
	}

	//Al crearse devuelve el _id y no el id

	c.IndentedJSON(http.StatusCreated, 1)
}
