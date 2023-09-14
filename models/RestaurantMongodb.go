package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (newRestaurant *RestaurantList) CreateRestaurantModel(ctx context.Context) (*mongo.InsertOneResult, error) {
	return RestaurantCollection.InsertOne(ctx, newRestaurant)
}

// func (UpdateRestaurant *RestaurantList) UpdateOneRestaurantsModel(ctx context.Context, *mongo.up) (*mongo.UpdateOneModel, error) {
// 	return RestaurantCollection.UpdateOne(ctx, UpdateRestaurant)
// }
