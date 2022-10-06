package controllers

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

func NewProject(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	var project models.Project

	if err := c.BindJSON(&project); err != nil { //traslada lo que tiene el contexto json a la variable golang user
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Como el autor es un campo requerido lo ingreso antes de validar.
	//session := sessions.Default(c)
	//user := session.Get("user")
	user := c.Value("username") // CREO
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
	} else {
		var userObj models.User
		err := userCollection.FindOne(ctx, bson.M{"username": user}).Decode(&userObj) //decodifica el json a golang luego de buscarlo en la tabla
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		project.Author = userObj.Username

		//validacion de los campos
		validationErr := validate.Struct(project) //compara y valida
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := projectCollection.CountDocuments(ctx, bson.M{"name": project.Name, "author": project.Author}) //Lo usamos para validar, si ya hay documentos con el mismo nombre y pertenezcan al mismo autor
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

func DeleteProject(c *gin.Context) { //SIN POBAR
	//Chequear si el usuario logeado es el autor del proyecto por las dudas
	project := c.Param("name")
	user := c.Value("username")
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	count, err := projectCollection.CountDocuments(ctx, bson.M{"name": project, "author": user}) //Lo usamos para validar, si ya hay documentos con el mismo nombre y pertenezcan al mismo autor
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

	/*fmt.Println(users) No devuelvo error porque si no hay documentos no tengo que romper la aplicacion
	if len(users) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No listed projects under specified user. Maybe it doesn't exist!"})
		return
	}*/

	c.JSON(http.StatusOK, users)
}

func GetExtProjects(c *gin.Context) {
	username := c.Param("username") //la idea es recuperar todos los projectos cuyo author id sea userId

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
	//fmt.Println(userprojects)
	if len(userprojects) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No listed projects under specified user. Maybe it doesn't exist!"})
		return
	}

	c.JSON(http.StatusOK, userprojects)

}

func GetProjects(c *gin.Context) {
	//session := sessions.Default(c)
	//user := session.Get("user")
	user := c.Value("username") // CREO
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

		cursor, err := projectCollection.Find(ctx, bson.M{"author": user}) //decodifica el json a golang luego de buscarlo en la tabla
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
