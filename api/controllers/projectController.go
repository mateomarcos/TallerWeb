package controllers

/*Controllers package is used to manage database connection for both models of our applications; User and Project.*/

import (
	"Portfolio/database"
	"Portfolio/models"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var projectCollection *mongo.Collection = database.OpenColllection(database.Client, "project")

/*
Input: Gin Context (Middleware to allow data flow and json requests/responses.), HTTP Request with project format.
Output: http response through gin context, database submission or error message.
*/
func NewProject(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	var project models.Project

	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.Value("username")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
	} else {
		var userObj models.User
		err := userCollection.FindOne(ctx, bson.M{"username": user}).Decode(&userObj)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		project.Author = userObj.Username

		validationErr := validate.Struct(project) //compara y valida
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := projectCollection.CountDocuments(ctx, bson.M{"name": project.Name, "author": project.Author})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking project name"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This project name already exists!"})
			return
		}

		project.ID = primitive.NewObjectID()
		project.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		resultInsertionNumber, insertError := projectCollection.InsertOne(ctx, project)
		if insertError != nil {
			msg := "Project item was not created"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

/*
Input: Gin Context (Middleware to allow data flow and json requests/responses.), parameter with project name to delete.
Output: http response through gin context, database deletion or error message.
*/
func DeleteProject(c *gin.Context) {
	project := c.Param("name")
	user := c.Value("username")
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	count, err := projectCollection.CountDocuments(ctx, bson.M{"name": project, "author": user})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking project name"})
		return
	}
	if count == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This project either doesn't belong to you or it doesn't exist!"})
		return
	} else {
		deleteResult, err := projectCollection.DeleteOne(ctx, bson.M{"name": project, "author": user})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while deleting project"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, deleteResult)
	}
}

/*
Input: Gin Context (Middleware to allow data flow and json requests/responses.)
Output: http response through gin context with the 4 most recent submitted projects.
*/
func GetActiveUsers(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

	options := options.Find()
	options.SetSort(bson.D{{"created_at", -1}})
	options.SetLimit(4)
	cursor, err := projectCollection.Find(ctx, bson.M{}, options)

	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var users []bson.M
	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, users)
}

/*
Input: Gin Context (Middleware to allow data flow and json requests/responses), username parameter.
Output: http response through gin context, error message or a list of projects submitted by given username.
*/
func GetExtProjects(c *gin.Context) {
	username := c.Param("username")

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

	cursor, err := projectCollection.Find(ctx, bson.M{"author": username})

	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var userprojects []bson.M
	if err = cursor.All(ctx, &userprojects); err != nil {
		log.Fatal(err)
	}
	if len(userprojects) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No listed projects under specified user. Maybe it doesn't exist!"})
		return
	}

	c.JSON(http.StatusOK, userprojects)

}

/*
Input: Gin Context (Middleware to allow data flow and json requests/responses).
Output: http response through gin context, error message or a list of projects submitted by authenticated user.
*/
func GetProjects(c *gin.Context) {
	user := c.Value("username")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

		cursor, err := projectCollection.Find(ctx, bson.M{"author": user})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var userprojects []bson.M
		if err = cursor.All(ctx, &userprojects); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, userprojects)
	}
}
