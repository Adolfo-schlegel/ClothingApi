package main

import (
	"os"

	BasicClient "example/src/handlers/Client/Basic"
	ParcialClient "example/src/handlers/Client/Parcial"
	User "example/src/handlers/User"
	"example/src/initializers"
	"example/src/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

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
	initializers.LoadEnvVariables()
	initializers.ConfigMongoConnection()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	host := "localhost:9990"
	router := gin.Default()
	router.Use(CORSMiddleware())
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

	//User
	goGroup.POST("/user", User.SignUp)
	goGroup.POST("/login", User.Login)
	goGroup.GET("/validate", middleware.RequireAuth, User.Validate)
	//Basic Client
	goGroup.GET("/clients", BasicClient.GetClients)
	goGroup.POST("/clients", BasicClient.CreateClient)
	goGroup.PATCH("/clients", BasicClient.ChangeRating)
	goGroup.GET("/clients/find/:id", BasicClient.GetclientById)
	goGroup.PUT("/clients/:id", BasicClient.UpdateClient)
	goGroup.DELETE("/clients/:id", BasicClient.DeleteById)

	//Partial Client
	goGroup.GET("/parcials", ParcialClient.GetPartials)
	goGroup.GET("/parcials/find/:id", ParcialClient.GetPartialById)
	goGroup.POST("/parcials", ParcialClient.CreateParcial)
	goGroup.DELETE("/parcials/:id", ParcialClient.DeleteById)

	//Partial BOXES
	goGroup.POST("/parcials/Box/:id", ParcialClient.CreateBox)
	goGroup.DELETE("/parcials/Box", ParcialClient.DeleteBox)
	goGroup.PUT("/parcials/Box", ParcialClient.Updatebox)
	goGroup.POST("/parcials/Box/Image", ParcialClient.AddImage)
	goGroup.DELETE("/parcials/Box/Image", ParcialClient.DeleteImage)

	router.Run(host)
}
