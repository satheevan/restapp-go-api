package UserController

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	c "github.com/pulsarcoder/Projects/restaurantgo/common"
	"github.com/pulsarcoder/Projects/restaurantgo/models"
	"github.com/pulsarcoder/Projects/restaurantgo/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// create function
func (uc UserController) UserCreate() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// user := models.Users{}
		var user = new(models.Users)
		defer cancel()
		//Define

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			// ============Original=========================
			rw.WriteHeader(http.StatusBadRequest)
			res := responses.UserResponse{Status: http.StatusBadRequest, Message: "Error not get from body request", Data: c.DataMap{
				"data": err.Error()}}
			json.NewEncoder(rw).Encode(res)
			return
		}
		// var validate = validator.New()

		//validate the request body(bodyparser)

		/*if validationErr := validate.Struct(&user); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			res := responses.UserResponse{Status: http.StatusBadRequest, Message: "Invalid data", Data: c.DataMap{
				"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(res)
			return
		}*/
		password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		newUser := models.Users{
			Id:        primitive.NewObjectID(),
			Name:      user.Name,
			Email:     user.Email,
			Password:  string(password),
			CreatedAt: time.Now(),
		}
		_, err := newUser.CreateUser(ctx)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			res := responses.UserResponse{Status: http.StatusInternalServerError, Message: "Error in Inserting into the database ", Data: c.DataMap{
				"error": "Internal server error"}}
			log.Println(err.Error())
			uc.Json(rw, res)
			return
		}
		//result
		rw.WriteHeader(http.StatusCreated)
		res := responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "successfully Inserting into the database ",
			Data:    c.DataMap{"data": newUser},
		}
		uc.Json(rw, res)

	} //return function end
} //function end
