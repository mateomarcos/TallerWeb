package main

import (
	"Portfolio/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
	router.Use(gin.Logger())
	router.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))))

	//Rutas base, no requieren autenticacion
	routes.AuthRoutes(router)

	private := router.Group("/user")
	private.Use(AuthRequired())
	routes.UserRoutes(private)
	routes.ProjectRoutes(private)

	router.Run(":" + port)

}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// You'd normally redirect to login page
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		} else {
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}
