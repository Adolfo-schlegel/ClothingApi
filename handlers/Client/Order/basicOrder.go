package BasicOrder

import (
	Model "example/src/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var ClientCol *mongo.Collection = nil
var OrderCol *mongo.Collection = nil

func CreateOrder(c *gin.Context) {
	idClient, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error al parsear id " + err.Error()})
	}
	var order Model.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	order.ID = primitive.NewObjectID()
	order.IdClient = idClient

	res, err := OrderCol.InsertOne(c, order)
	fmt.Printf("res: %v\n", res)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error al insertar Usuario en la base de datos " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, order.ID)
}
func GetOrders(c *gin.Context) {
	cursor, err := OrderCol.Find(c, bson.M{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"\nmessage": "Error parsing JSON", "code": http.StatusInternalServerError})
		return
	}

	defer cursor.Close(c)

	var results []Model.Order

	for cursor.Next(c) {
		var result Model.Order
		err := cursor.Decode(&result)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"\nmessage": "Error parsing JSON: " + err.Error(), "code": http.StatusInternalServerError})
			return
		}
		results = append(results, result)
	}

	c.IndentedJSON(http.StatusOK, results)
}
