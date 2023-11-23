package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"

	MongoDB "example/src/DB/MongoDB"
	BasicClient "example/src/handlers/Client/Basic"
	ParcialClient "example/src/handlers/Client/Parcial"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// The init function will run before our main function to establish a connection to MongoDB.
const uri = "mongodb://dolphin:10deagostO@192.168.0.130:27017"

func ConfigMongoConnection() {
	fmt.Print("------------Connection to MongoDB------------:\n ")
	fmt.Print(uri + "\n")

	if err := MongoDB.Connect_to_mongodb(uri); err != nil {
		log.Fatal("Could not connect to MongoDB\n" + err.Error())
	}
	dt := MongoDB.MongoClient.Database("Confecciones")

	//Clientes
	BasicClient.ClientCol = dt.Collection("clients")
	BasicClient.ParcialCol = dt.Collection("PartialClients")
	ParcialClient.Collection = dt.Collection("PartialClients")
	//IntegralClient.Collection = dt.Collection("integralClients")

	//Prendas

	//Diseños

	//Facturacion
}
func verifyOutput() {
	//Verificar si os.Stdout está conectado a una terminal
	if fi, err := os.Stdout.Stat(); err == nil {
		if (fi.Mode() & os.ModeCharDevice) == 0 {
			// La salida está redirigida (por ejemplo, a un archivo o un proceso)
			fmt.Println("La salida está redirigida.")
		} else {
			// La salida está conectada a una terminal
			fmt.Println("La salida está conectada a una terminal.")
		}
	} else {
		fmt.Println("Error al obtener información de os.Stdout:", err)
	}
}
func init() {
	verifyOutput()
	ConfigMongoConnection()
}

func main() {
	host := "localhost:9990"
	router := gin.Default()
	print("Router\n")

	// Crear un grupo de rutas con el prefijo /api/go
	goGroup := router.Group("/api/go")

	goGroup.GET("/", func(c *gin.Context) {
		c.String(200, "Bienvenido a go")
	})

	goGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	goGroup.GET("/health", func(c *gin.Context) {
		c.String(200, "Healthy")
	})

	//--------------------------------------------Middleware ----------------------------------------------------------------
	//router.Use(cors.Default())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	//----------------------------------------------------------------------------------------------------------------------------

	//Basic Client Info
	goGroup.GET("/clients", BasicClient.GetClients)
	goGroup.POST("/clients", BasicClient.CreateClient)
	goGroup.PATCH("/clients", BasicClient.ChangeRating)
	goGroup.GET("/clients/find/:id", BasicClient.GetclientById)
	goGroup.PUT("/clients/:id", BasicClient.UpdateClient)
	goGroup.DELETE("/clients/:id", BasicClient.DeleteById)

	//Parcial Client
	goGroup.GET("/parcials", ParcialClient.GetPartials)
	goGroup.GET("/parcials/find/:id", ParcialClient.GetPartialById)
	goGroup.POST("/parcials", ParcialClient.CreateParcial)

	router.Run(host)
}
