package RestaurantController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pulsarcoder/Projects/restaurantgo/common"
	"github.com/pulsarcoder/Projects/restaurantgo/models"
	"github.com/pulsarcoder/Projects/restaurantgo/requests"
	"github.com/pulsarcoder/Projects/restaurantgo/responses"
)

func (rc RestaurantController) CreateRestaurants() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// user := models.Users{}
		//create
		var userPayload = new(requests.RestaurantCreateRequest)
		defer cancel()

		//Getting data from payload
		if err := json.NewDecoder(r.Body).Decode(&userPayload); err != nil {
			fmt.Print("Error in getting in payload values", err.Error())
			// ============Original=========================
			rw.WriteHeader(http.StatusBadRequest)
			res := responses.UserResponse{Status: http.StatusBadRequest, Message: "Error not get from body request", Data: common.DataMap{
				"data": err.Error()}}
			json.NewEncoder(rw).Encode(res)
			return
		}
		// assigning
		newRestaurant := models.RestaurantList{
			Name:      userPayload.Name,
			Contact:   userPayload.Contact,
			Address:   userPayload.Address,
			CreatedAt: time.Now(),
		}
		//store in database
		_, err := newRestaurant.CreateRestaurantModel(ctx)
		if err != nil {
			fmt.Print("Error in database and the value not saved in mongodb", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			res := responses.UserResponse{Status: http.StatusInternalServerError, Message: "Error in Inserting into the database ", Data: common.DataMap{
				"error": "Internal server error"}}
			log.Println(err.Error())
			rc.Json(rw, res)
			return
		}
		//result
		rw.WriteHeader(http.StatusCreated)
		res := responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "successfully Inserting into the database ",
			Data:    common.DataMap{"data": newRestaurant},
		}
		rc.Json(rw, res)

	}
}
