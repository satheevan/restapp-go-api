package RestaurantController

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pulsarcoder/Projects/restaurantgo/common"
	"github.com/pulsarcoder/Projects/restaurantgo/models"
	"github.com/pulsarcoder/Projects/restaurantgo/requests"
	"github.com/pulsarcoder/Projects/restaurantgo/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (rc RestaurantController) UpdateOneRestaurants() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		// create
		var userPayload = new(requests.RestaurantCreateRequest)

		//getting the Query parameter for filiter "id"
		routeParams := mux.Vars(r)
		RestaurantId := routeParams["id"]

		defer cancel()

		ObjectId, _ := primitive.ObjectIDFromHex(RestaurantId)
		filter := bson.M{"_id": ObjectId}

		if err := json.NewDecoder(r.Body).Decode(&userPayload); err != nil {
			fmt.Print("RestaurantController : UpdateRestaurantController :UpdateOneRestaurants => Error in getting data from payLoad", err.Error())
			//fornt end
			rw.WriteHeader(http.StatusBadRequest)
			res := responses.UserResponse{Status: http.StatusBadRequest, Message: "Error not get from body request", Data: common.DataMap{
				"data": err.Error()}}
			json.NewEncoder(rw).Encode(res)
		}
		//assigning
		UpdateRestaurant := bson.M{
			"name":       userPayload.Name,
			"address":    userPayload.Address,
			"contact":    userPayload.Contact,
			"modifiedAt": time.Now(),
		}
		// Name =       userPayload.Name,
		// Contact:    userPayload.Contact,
		// Address:    userPayload.Address,
		// ModifiedAt: time.Now(),

		// UpdateRestaurant["Name"] = userPayload.Name,
		// UpdateRestaurant["Contact"] = userPayload.Contact,
		// UpdateRestaurant["Address"] = userPayload.Address,
		// UpdateRestaurant["ModifiedAt"] = time.Now(),

		update := bson.M{"$set": UpdateRestaurant}

		result, err := models.RestaurantCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			fmt.Println("Restaurantcontroller : UpdsteRestaurantController: UpdateOneRestaurants=> Error in Getting the data from database:", err)
			res := rc.Responses(http.StatusInternalServerError, "Error in data receives", common.DataMap{"Error": err.Error()})
			rc.Json(rw, res)
			return
		}
		rw.WriteHeader(http.StatusOK)
		res := rc.Responses(http.StatusOK, "successfully getting", common.DataMap{"data": result})
		rc.Json(rw, res)
	}
}
