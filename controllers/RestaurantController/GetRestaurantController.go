package RestaurantController

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pulsarcoder/Projects/restaurantgo/common"
	"github.com/pulsarcoder/Projects/restaurantgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (rc RestaurantController) GetAllRestaurants() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//created
		var restaurant []models.RestaurantList
		defer cancel()

		//database query
		result, err := models.RestaurantCollection.Find(ctx, bson.M{})

		if err != nil {
			log.Fatal("Restaurantcontroller : GetRestaurantController: GetAllRestaurants=> Error in database Not fetching the data")
			rc.Status500(rw)
			res := rc.Responses(http.StatusInternalServerError, "Error in data receives", common.DataMap{"Error": err.Error()})
			rc.Json(rw, res)
			return
		}
		// reading from DB in optimal way
		defer result.Close(ctx)

		for result.Next(ctx) {
			var singleRestaurant models.RestaurantList
			if err = result.Decode(&singleRestaurant); err != nil {
				fmt.Println("Restaurantcontroller : GetRestaurantController: GetAllRestaurants=> Error in Getting the data from database")
				rw.WriteHeader(http.StatusInternalServerError)
				res := rc.Responses(http.StatusInternalServerError, "Error in data receives", common.DataMap{"Error": err.Error()})
				rc.Json(rw, res)
				return
			}
			restaurant = append(restaurant, singleRestaurant)
		}
		rw.WriteHeader(http.StatusOK)
		res := rc.Responses(http.StatusOK, "data receives successfully", common.DataMap{"data": restaurant})
		rc.Json(rw, res)
	}
}

// Get One data
func (rc RestaurantController) GetOneRestaurants() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		// create
		var restaurant = new(models.RestaurantList)

		//getting the Query parameter
		routeParams := mux.Vars(r)
		RestaurantId := routeParams["Id"]

		defer cancel()

		ObjectId, _ := primitive.ObjectIDFromHex(RestaurantId)
		if err := models.RestaurantCollection.FindOne(ctx, bson.M{"Id": ObjectId}).Decode(&restaurant); err != nil {
			fmt.Println("Restaurantcontroller : GetRestaurantController: GetOneRestaurants=> Error in Getting the data from database")
			res := rc.Responses(http.StatusInternalServerError, "Error in data receives", common.DataMap{"Error": err.Error()})
			rc.Json(rw, res)
		}
		rw.WriteHeader(http.StatusOK)
		res := rc.Responses(http.StatusOK, "successfully getting", common.DataMap{"data": restaurant})
		rc.Json(rw, res)
	}
}
