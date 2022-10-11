package controllers

/*Controllers package is used to manage database connection for both models of our applications; User and Project.*/

import (
	"Portfolio/database"
	"Portfolio/models"
	"context"
	"log"
	"os"
	"time"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenColllection(database.Client, "user")
var validate = validator.New()
var SECRET_KEY string = os.Getenv("SECRET_KEY")

/*
Input: Gin Context (Middleware to allow data flow and json requests/responses.), HTTP Request with user format.
Output: http response through gin context, database submission or error message.
*/
func Signup(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

	//Verify if user with given username already exists
	count, err := userCollection.CountDocuments(ctx, bson.M{"username": user.Username})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking email"})
		return
	}

	password := HashPassword(*user.Password)
	user.Password = &password

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
	c.JSON(http.StatusOK, resultInsertionNumber)
}

/*
Input: password string.
Output: encrypted string
*/
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

/*
Input: username string.
Output: jwt tokens that contains expiration and username.
*/
func generateJWT(username string) (string, error) {

	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(4 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*
Input: Gin Context (Middleware to allow data flow and json requests/responses.), HTTP Request with user format.
Output: http response through gin context with jwt token or error message.
*/
func Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

	var foundUser models.User
	err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
	if !passwordIsValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, err := generateJWT(*user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user", "token": token})
}

/*
Input: Database retrieved password string, provided password string.
Output: boolean and message wheter the encrypted passwords match or not.
*/
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := "Incorrect password!"

	if err != nil {
		check = false
	}
	return check, msg
}

/* Not used */
func GetUser(c *gin.Context) {

	user := c.Value("username")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

		var retUser models.User
		err := userCollection.FindOne(ctx, bson.M{"username": user}).Decode(&retUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, retUser)
	}
}
