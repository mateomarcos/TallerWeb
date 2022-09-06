package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct { //middleware entre programa y la base de datos mongo, practicamente golang a json y json a golang
	ID              primitive.ObjectID `bson:"_id`
	Name            *string            `json:"name" validate:"required,min=5,max=100"`
	Description     *string            `json:"description" validate:"required,min=2,max=255"`
	Repository_link *string            `json:"repository_link"`
	Created_at      time.Time          `json:"created_at"`
	Author_uid      *string            `json:"author_uid" validate:"required,min=1"`
}
