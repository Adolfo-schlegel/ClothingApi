package MongoDB

import (
	//Go packages
	"context"

	// Add the MongoDB driver packages
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

//func GetCollection(collectionName string, database string) *mongo.Collection {
//	db := MongoClient.Database(database)
//	collection := db.Collection(collectionName)
//	return collection
//}

// Connecting to MongoDB
func Connect_to_mongodb(uri string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		fmt.Print("Seteando datos del cliente")
		fmt.Print(err)
		return err
	}

	//check conn
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Print("Error al verificar la conexion\n")
		fmt.Print(err)
		return err
	}

	MongoClient = client

	return err
}
