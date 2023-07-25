package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (newUser *Users) CreateUser(ctx context.Context) (*mongo.InsertOneResult, error) {
	newUser.Id = primitive.NewObjectID()
	return UserCollection.InsertOne(ctx, newUser)
}
