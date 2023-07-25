package models

import (
	"time"

	"github.com/pulsarcoder/Projects/restaurantgo/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RestaurantList struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Name       string             `json:"name,omitempty" validate:"required"`
	Contact    string             `json:"contact,omitempty" validate:"required"`
	Address    string             `json:"address,omitempty" validate:"required"`
	CreatedAt  time.Time          `bson:"created_at" json:"createdAt,omitempty"`
	ModifiedAt time.Time          `bson:"modifiedAt" json:"modifiedAt,omitempty"`
	DeletedAt  time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}

var RestaurantCollection *mongo.Collection = configs.GetCollection(configs.DB, "restaurantlists")
