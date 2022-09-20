package main

import (
	"Portfolio/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	//router.Use(gin.Logger())
	//router.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))))

	//Rutas base, no requieren autenticacion
	routes.AuthRoutes(router)

	private := router.Group("/user")
	private.Use(AuthRequired())
	routes.UserRoutes(private)
	routes.ProjectRoutes(private)

	router.Run(":" + port)

}

// "MIDDLEWARE"
func ValidateToken(signedToken string) (claims jwt.MapClaims, msg string) {
	token, err := jwt.ParseWithClaims(signedToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(SECRET_KEY)), nil
	},
	)
	//fmt.Println("TOKEN: ", token)
	/*fmt.Println("ERROR: ", err) Probando con THunderclient y Postman dice que signature es invalida, si solo pongo el token ahic
	if err != nil {
		msg = err.Error()
		return nil, msg
	}*/
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)

	if !ok {
		msg = fmt.Sprintf("The token is invalid")
		msg = err.Error()
		return nil, msg
	}

	if int64(claims["exp"].(float64)) < time.Now().Local().Unix() {
		msg = fmt.Sprint("The token has expired")
		msg = err.Error()
		return
	}

	return claims, msg
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err := ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("username", claims["username"])
		c.Set("authorized", claims["authorized"])
		c.Next()
	}
}
