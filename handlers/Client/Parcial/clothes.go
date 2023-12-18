package ParcialClient

import (
	Model "example/src/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBox(c *gin.Context) {
	idClient, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error al parsear id " + err.Error()})

	}

	var box Model.Box

	//get the client from the json request
	if err := c.ShouldBindJSON(&box); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Set a new ObjectID for the client
	box.ID = primitive.NewObjectID()

	// Update the document by adding the new item to the clothes array
	update := bson.D{{"$push", bson.D{{"clothes", box}}}}

	//Insert the client into the database
	_, err = Collection.UpdateOne(c, bson.M{"idClient": idClient}, update)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, box.ID)
}
func AddImage(c *gin.Context) {
	idcli, ok := c.GetQuery("idClient")
	idObj, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id or rating query parameter."})
	}

	var image Model.Image

	if err := c.ShouldBindJSON(&image); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body" + err.Error()})
	}

	idClient, err := primitive.ObjectIDFromHex(idcli)
	id, err := primitive.ObjectIDFromHex(idObj)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error al parsear id " + err.Error()})
		return
	}

	// Define un filtro para el objeto que quieres eliminar.
	filter := bson.M{
		"idClient": idClient,
		"clothes": bson.M{
			"$elemMatch": bson.M{"_id": id},
		},
	}

	// Update the document by adding the new item to the clothes array
	update := bson.D{{"$push", bson.D{{"clothes.$.embroideryImages", image}}}}

	//Update the document in the database
	_, err = Collection.UpdateOne(c, filter, update)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error: " + err.Error()})
		return
	}

}

func DeleteImage(c *gin.Context) {
	idcli, ok := c.GetQuery("idClient")
	idObj, ok := c.GetQuery("id")
	nameFile, ok := c.GetQuery("nameFile")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id or rating query parameter."})
	}

	idClient, err := primitive.ObjectIDFromHex(idcli)
	id, err := primitive.ObjectIDFromHex(idObj)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error al parsear id " + err.Error()})
		return
	}

	// Define un filtro para el objeto que quieres eliminar.
	filter := bson.M{
		"idClient": idClient,
		"clothes": bson.M{
			"$elemMatch": bson.M{"_id": id},
		},
	}

	// Update the document by adding the new item to the clothes array
	update := bson.D{{"$pull", bson.D{{"clothes.$.embroideryImages", bson.D{{"name", nameFile}}}}}}

	//Update the document in the database
	_, err = Collection.UpdateOne(c, filter, update)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error: " + err.Error()})
		return
	}

}

func DeleteBox(c *gin.Context) {
	idcli, ok := c.GetQuery("idClient")
	idObj, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id or rating query parameter."})
	}

	idClient, err := primitive.ObjectIDFromHex(idcli)
	id, err := primitive.ObjectIDFromHex(idObj)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error al parsear id " + err.Error()})
		return
	}

	// Define un filtro para el objeto que quieres eliminar.
	filter := bson.M{"idClient": idClient}

	// Update the document by removing the specific item from the clothes array
	update := bson.D{{"$pull", bson.D{{"clothes", bson.D{{"_id", id}}}}}}

	//Update the document in the database
	_, err = Collection.UpdateOne(c, filter, update)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, 1)
}
func Updatebox(c *gin.Context) {

	idcli, ok := c.GetQuery("idClient")
	idobj, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id or rating query parameter."})
	}

	idClient, err := primitive.ObjectIDFromHex(idcli)
	id, err := primitive.ObjectIDFromHex(idobj)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body" + err.Error()})
	}

	var box Model.Box

	if err := c.ShouldBindJSON(&box); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body" + err.Error()})
	}

	box.ID = id
	filter := bson.D{{"idClient", idClient}, {"clothes._id", id}}
	update := bson.M{"$set": bson.D{{"clothes.$", box}}}
	result, err := Collection.UpdateOne(c, filter, update)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error: " + err.Error()})
	}
	c.IndentedJSON(http.StatusOK, result)
}
