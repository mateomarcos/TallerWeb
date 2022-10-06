package main

import (
	"Portfolio/controllers"
	"Portfolio/routes"
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

	router := gin.New()
	router.Use(CORSMiddleware())
	router.GET("/activeUsers", controllers.GetActiveUsers)

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
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		//msg = fmt.Sprintf("The token is invalid")
		msg = err.Error()
		return nil, msg
	}

	if int64(claims["exp"].(float64)) < time.Now().Local().Unix() {
		//msg = fmt.Sprint("The token has expired")
		msg = err.Error()
		return
	}

	return claims, msg
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Token header provided"})
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		//fmt.Println("REQUEST METHOD: ", c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			//fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
