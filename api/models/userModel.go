package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Middleware between the database and mongo driver, in which each field is linked to a column in the User collection of mongo.
*/
type User struct {
	ID       primitive.ObjectID `bson:"_id`
	Username *string            `json:"username" validate:"required,min=2,max=100"`
	Password *string            `json:"password" validate:"required,min=6"`
}
