package initializers

import (
	"example/src/DB/MongoDB"
	BasicClient "example/src/handlers/Client/Basic"
	BasicOrder "example/src/handlers/Client/Order"
	ParcialClient "example/src/handlers/Client/Parcial"
	User "example/src/handlers/User"
	"example/src/middleware"
	"fmt"
	"log"
	"os"
)

func ConfigMongoConnection() {
	fmt.Print("------------Connection to MongoDB------------:\n ")
	db := os.Getenv("MONGODB")

	fmt.Print(db + "\n")
	if err := MongoDB.Connect_to_mongodb(db); err != nil {
		log.Fatal("Could not connect to MongoDB\n" + err.Error())
	}
	dt := MongoDB.MongoClient.Database("Confecciones")

	//User
	User.UserCol = dt.Collection("Users")
	middleware.UserCol = dt.Collection("Users")

	//Clientes
	BasicClient.ClientCol = dt.Collection("clients")
	BasicClient.ParcialCol = dt.Collection("PartialClients")
	ParcialClient.Collection = dt.Collection("PartialClients")
	BasicOrder.OrderCol = dt.Collection("Orders")

	//IntegralClient.Collection = dt.Collection("integralClients")

	//Prendas

	//Diseños

	//Facturacion
}
