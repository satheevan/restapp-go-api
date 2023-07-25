package RestaurantController

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pulsarcoder/Projects/restaurantgo/common"
	"github.com/pulsarcoder/Projects/restaurantgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (rc RestaurantController) UpdateOneRestaurants() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		// create
		var restaurant = new(models.RestaurantList)

		//getting the Query parameter
		routeParams := mux.Vars(r)
		RestaurantId := routeParams["id"]
		// var restaurant = new(models.RestaurantList)

		defer cancel()

		ObjectId, _ := primitive.ObjectIDFromHex(RestaurantId)
		if err := models.RestaurantCollection.FindOne(ctx, bson.M{"id": ObjectId}).Decode(&restaurant); err != nil {
			fmt.Println("Restaurantcontroller : UpdsteRestaurantController: UpdateOneRestaurants=> Error in Getting the data from database")
			res := rc.Responses(http.StatusInternalServerError, "Error in data receives", common.DataMap{"Error": err.Error()})
			rc.Json(rw, res)
			return
		}
		rw.WriteHeader(http.StatusOK)
		res := rc.Responses(http.StatusOK, "successfully getting", common.DataMap{"data": restaurant})
		rc.Json(rw, res)
	}
}
