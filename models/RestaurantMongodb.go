package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (newRestaurant *RestaurantList) CreateRestaurantModel(ctx context.Context) (*mongo.InsertOneResult, error) {
	newRestaurant.Id = primitive.NewObjectID()
	return RestaurantCollection.InsertOne(ctx, newRestaurant)
}
