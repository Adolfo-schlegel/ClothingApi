package middleware

import (
	Model "example/src/Models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCol *mongo.Collection = nil

func RequireAuth(c *gin.Context) {
	fmt.Println("In middleware")

	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//check expiration time
		if float64(time.Now().Unix()) > claims["expires"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//check if the claim's user is the same as the database user
		var user Model.User

		idUser, err := primitive.ObjectIDFromHex(claims["id"].(string))

		filter := bson.D{{"_id", idUser}}

		err = UserCol.FindOne(c.Request.Context(), filter).Decode(&user)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
