package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct { //middleware entre programa y la base de datos mongo, practicamente golang a json y json a golang
	ID       primitive.ObjectID `bson:"_id`
	Username *string            `json:"username" validate:"required,min=2,max=100"`
	Password *string            `json:"password" validate:"required,min=6"`
}
