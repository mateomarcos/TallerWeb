package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Middleware between the database and mongo driver, in which each field is linked to a column in the Project collection of mongo.
*/
type Project struct {
	ID          primitive.ObjectID `bson:"_id`
	Name        *string            `json:"name" validate:"required,min=5,max=100"`
	Description *string            `json:"description" validate:"required,min=2,max=255"`
	Repository  *string            `json:"repository"`
	Created_at  time.Time          `json:"created_at"`
	Author      *string            `json:"author" validate:"required,min=1"`
}
