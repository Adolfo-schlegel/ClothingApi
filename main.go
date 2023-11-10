package main

import (
	MongoDB "example/src/DB/MongoDB"
	Client "example/src/handlers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// The init function will run before our main function to establish a connection to MongoDB.
const uri = "mongodb://dolphin:10deagostO@192.168.0.130:27017"

func init() {

	fmt.Print("------------Connection to MongoDB------------:\n ")
	fmt.Print(uri + "\n")

	if err := MongoDB.Connect_to_mongodb(uri); err != nil {
		log.Fatal("Could not connect to MongoDB\n" + err.Error())
	}
	dt := MongoDB.MongoClient.Database("Confecciones")

	//Clientes
	Client.Collection = dt.Collection("clients")

	//Prendas

	//Facturacion
}

func main() {
	router := gin.Default()
	print("Router\n")

	//Client
	router.GET("/clients/:id", Client.GetclientById)
	router.POST("/clients/:id", Client.UpdateClient)
	router.DELETE("/clients/:id", Client.DeleteById)
	router.POST("/clients/Create", Client.CreateClient)
	router.GET("/clients/All", Client.GetClients)
	router.PATCH("/clients/changeRating", Client.ChangeRating)

	router.Run("localhost:8080")
}
