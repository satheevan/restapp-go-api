package models

import (
	"time"

	"github.com/pulsarcoder/Projects/restaurantgo/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// signup
type Users struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Name       string             `json:"name,omitempty" validate:"required"`
	Email      string             `json:"email,omitempty" validate:"required"`
	Password   string             `json:"password,omitempty" validate:"required"`
	CreatedAt  time.Time          `bson:"created_at" json:"createdAt,omitempty"`
	ModifiedAt time.Time          `bson:"modifiedAt" json:"modifiedAt,omitempty"`
	DeletedAt  time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}

var UserCollection *mongo.Collection = configs.GetCollection(configs.DB, "registers")

// func CreateUser(ctx context.Context, newUser Users) Users, error {
// 	result, error := UserCollection.InsertOne(ctx, newUser)
// }
