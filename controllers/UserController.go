package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	// folders
	"github.com/gorilla/mux"
	"github.com/pulsarcoder/Projects/restaurantgo/configs"
	"github.com/pulsarcoder/Projects/restaurantgo/models"
	"github.com/pulsarcoder/Projects/restaurantgo/responses"
)

var UserCollection *mongo.Collection = configs.GetCollection(configs.DB, "registers")

type UserController struct {
	BaseController
}

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
			res := responses.UserResponse{Status: http.StatusBadRequest, Message: "Error not get from body request", Data: map[string]interface{}{
				"data": err.Error()}}
			json.NewEncoder(rw).Encode(res)
			return
		}
		var validate = validator.New()

		//validate the request body(bodyparser)

		if validationErr := validate.Struct(&user); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			res := responses.UserResponse{Status: http.StatusBadRequest, Message: "Invalid data", Data: map[string]interface{}{
				"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(res)
			return
		}
		password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		newUser := models.Users{
			Id:       primitive.NewObjectID(),
			Name:     user.Name,
			Email:    user.Email,
			Password: string(password),
		}
		_, err := UserCollection.InsertOne(ctx, newUser)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			res := responses.UserResponse{Status: http.StatusInternalServerError, Message: "Error in Inserting into the database ", Data: map[string]interface{}{
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
			Data:    map[string]interface{}{"data": newUser},
		}
		uc.Json(rw, res)

	} //return function end
} //function end

// get data
// url: http://localhost:<port>/getuser/{email}  => FOR GET METHOD query parrams
func (uc UserController) UserGetdata() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		routeParams := mux.Vars(r)
		UserEmail := routeParams["email"]
		var users = new(models.Users)

		defer cancel()
		// ObjectId,_ :=primitive.ObjectIDFromHex(UserID)
		if err := UserCollection.FindOne(ctx, bson.M{"email": UserEmail}).Decode(&users); err != nil {
			res := responses.UserResponse{Status: http.StatusInternalServerError, Message: "Error in Inserting into the database ", Data: map[string]interface{}{
				"error": "Internal server error"}}
			log.Println(err.Error())
			uc.Json(rw, res)
			return
		}
		rw.WriteHeader(http.StatusOK)
		res := uc.Responses(http.StatusOK, "successfully getting", map[string]interface{}{"data": users})
		uc.Json(rw, res)
	}

}
