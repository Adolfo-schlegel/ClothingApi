package BasicClient

import (
	Model "example/src/Models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ClientCol *mongo.Collection = nil
var ParcialCol *mongo.Collection = nil

func GetclientById(c *gin.Context) {
	client, err := findClient(c.Param("id"), c)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"\nmessage": "Client not found." + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, client)
}

func findClient(clientId string, c *gin.Context) (*Model.Client, error) {
	var result Model.Client

	// Create a primitive.ObjectID from a hexadecimal string
	id, err := primitive.ObjectIDFromHex(clientId)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{"id", id}}

	err = ClientCol.FindOne(c.Request.Context(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func intToHexObjectID(value int) (primitive.ObjectID, error) {
	// Convert the integer to a hexadecimal string
	hexString := fmt.Sprintf("%x", value)

	// Create a primitive.ObjectID from the hexadecimal string
	objectID, err := primitive.ObjectIDFromHex(hexString)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return objectID, nil
}

func GetClients(c *gin.Context) {
	//Find all
	cursor, err := ClientCol.Find(c, bson.M{})

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"\nmessage": "Error parsing JSON", "code": http.StatusInternalServerError})
		return
	}

	defer cursor.Close(c)

	var results []Model.Client

	for cursor.Next(c) {
		var result Model.Client
		err := cursor.Decode(&result)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"\nmessage": "Error parsing JSON: " + err.Error(), "code": http.StatusInternalServerError})
			return
		}
		results = append(results, result)
	}

	c.IndentedJSON(http.StatusOK, results)
}

func CreateClient(c *gin.Context) {

	//get the client from the json request
	var client Model.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Set a new ObjectID for the client
	client.ID = primitive.NewObjectID()

	//Insert the client into the database
	_, err := ClientCol.InsertOne(c, client)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error al insertar Usuario en la base de datos " + err.Error()})
		return
	}

	if client.Type == "Parcial" {
		var parcial Model.ParcialClient

		// Setting default partial client
		parcial.ID = primitive.NewObjectID()
		parcial.IdClient = client.ID.Hex()
		parcial.Clothes = &[]Model.Clothe{}

		//Insert the partial client into the database
		_, err := ParcialCol.InsertOne(c, parcial)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "No se pudo crear cliente parcial " + err.Error()})
			return
		}

	} else if client.Type == "Integral" {
		//Insertar Integral

	} else if client.Type == "Nuevo" {
		//??
	}

	//fijarse que onda con el id insertado

	c.IndentedJSON(http.StatusCreated, 1)
}

func DeleteById(c *gin.Context) {

	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	result, err := ClientCol.DeleteOne(c, bson.M{"id": objectId})

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error: " + err.Error()})
	}

	c.IndentedJSON(http.StatusOK, result.DeletedCount)
}

func ChangeRating(c *gin.Context) {
	id, ok := c.GetQuery("id")
	ratingStr, ok := c.GetQuery("rating")
	rating, err := strconv.Atoi(ratingStr)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id or rating query parameter."})
	}

	client, err := findClient(id, c)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Client not found."})
		return
	}

	client.Rating = rating

	fmt.Println(client)

	filter := bson.D{{"id", client.ID}}
	update := bson.M{"$set": client}

	result, err := ClientCol.UpdateOne(c, filter, update)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func UpdateClient(c *gin.Context) {
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body" + err.Error()})
	}

	var client Model.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body" + err.Error()})
	}

	filter := bson.D{{"id", objectId}}
	update := bson.M{"$set": client}

	result, err := ClientCol.UpdateOne(c, filter, update)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error: " + err.Error()})
	}
	c.IndentedJSON(http.StatusOK, result)
}
