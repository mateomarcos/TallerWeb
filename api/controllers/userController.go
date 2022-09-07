package controllers

import (
	"Portfolio/database"
	"Portfolio/models"
	"context"
	"log"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/contrib/sessions"
)

var userCollection *mongo.Collection = database.OpenColllection(database.Client, "user")
var validate = validator.New()

func Signup(c *gin.Context) {
	//username := c.PostForm("username")
	//password := c.PostForm("password")

	var user models.User
	if err := c.BindJSON(&user); err != nil { //traslada lo que tiene el contexto json a la variable golang user
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(user) //compara y valida los parametros
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

	//Verify if user with given username already exists
	count, err := userCollection.CountDocuments(ctx, bson.M{"username": user.Username}) //Lo usamos para validar, si ya hay documentos con el mismo mail
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking email"})
		return
	}

	password := HashPassword(*user.Password)
	user.Password = &password
	// podriamos repetir lo de count  y err para otros atributos

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This username already exists!"})
		return
	}

	user.ID = primitive.NewObjectID()

	resultInsertionNumber, insertError := userCollection.InsertOne(ctx, user)
	defer cancel()
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User item was not created"})
		return
	}
	//defer cancel()
	c.JSON(http.StatusOK, resultInsertionNumber)
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	var user models.User
	if err := c.BindJSON(&user); err != nil { //traslada lo que tiene el contexto json a la variable golang user
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(user) //compara y valida los parametros
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

	// Check for username and password match, from Mongo to User middelware to hashpassword match
	var foundUser models.User
	err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser) //decodifica el json a golang luego de buscarlo en la tabla
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Password check
	passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
	if !passwordIsValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	// Save the username in the session
	session.Set("user", user.Username) // In real world usage you'd set this to the users ID ~ PENDIENTE CON ID NO FUNCIONA
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		check = false
	}
	return check, msg
}

func GetUser(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100) //REVISAR CANCEL

		var retUser models.User
		err := userCollection.FindOne(ctx, bson.M{"username": user}).Decode(&retUser) //decodifica el json a golang luego de buscarlo en la tabla
		defer cancel()                                                                //DUDA
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, retUser)
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete("user")
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
