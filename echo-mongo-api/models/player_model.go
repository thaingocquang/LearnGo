package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Player ...
type Player struct {
	Id primitive.ObjectID `json:"id,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
	Age int `json:"age,omitempty" validate:"required"`
	CreateAt time.Time
}
