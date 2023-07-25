package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	// folders
	"github.com/gorilla/mux"
	c "github.com/pulsarcoder/Projects/restaurantgo/common"
	"github.com/pulsarcoder/Projects/restaurantgo/configs"
	"github.com/pulsarcoder/Projects/restaurantgo/models"
	"github.com/pulsarcoder/Projects/restaurantgo/requests"
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
		var userPayload = new(requests.UserCreateRequest)
		defer cancel()
		//Define

		if err := json.NewDecoder(r.Body).Decode(&userPayload); err != nil {
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
		encryptedPwd := userPayload.EncryptPassword()
		newUser := models.Users{
			Name:      userPayload.Name,
			Email:     userPayload.Email,
			Password:  encryptedPwd,
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

// get data
// url: http://localhost:<port>/getuser/{email}  => FOR GET METHOD
func (uc UserController) UserGetdata() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// both getting the query params
		routeParams := mux.Vars(r)
		UserEmail := routeParams["email"]
		var users = new(models.Users)

		defer cancel()
		// ObjectId,_ :=primitive.ObjectIDFromHex(UserID)
		if err := UserCollection.FindOne(ctx, bson.M{"email": UserEmail}).Decode(&users); err != nil {
			res := responses.UserResponse{Status: http.StatusInternalServerError, Message: "Error in Inserting into the database ", Data: c.DataMap{
				"error": "Internal server error"}}
			log.Println(err.Error())
			uc.Json(rw, res)
			return
		}
		rw.WriteHeader(http.StatusOK)
		res := uc.Responses(http.StatusOK, "successfully getting", c.DataMap{"data": users})
		uc.Json(rw, res)
	}

}

// login details
func (uc UserController) LoginUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var userPayload = requests.UserCreateRequest{}
		// UserLogin := user.Email, user.Password
		var User models.Users
		defer cancel()
		// bodyrequest
		if err := json.NewDecoder(r.Body).Decode(&userPayload); err != nil {
			fmt.Println("UserController.LoginUser: Error in unmarshalling request body")
			uc.Status400(rw)
			res := uc.Responses(http.StatusBadRequest, "Invalid username or password", c.DataMap{"error": err.Error()})
			uc.Json(rw, res)
			return
		}
		//validate if nil
		if userPayload.Email == "" || userPayload.Password == "" {
			rw.WriteHeader(http.StatusBadRequest)
			uc.Json(rw, uc.Responses(http.StatusBadRequest, "Invalid username or password", c.DataMap{"error": "Invalid username or password"}))
			return
		}
		// query code => mongodb
		if err := UserCollection.FindOne(ctx, bson.M{"email": userPayload.Email}).Decode(&User); err != nil {
			fmt.Println("UserController.LoginUser: Error in the value comparring in DataBase", userPayload)
			uc.Status400(rw)
			res := uc.Responses(http.StatusBadRequest, "Incorrect Username or password", c.DataMap{"error": err.Error()})
			uc.Json(rw, res)
			return
		}
		// password comparison
		if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(userPayload.Password)); err != nil {
			fmt.Println("UserController.LoginUser: Error in comparing password using bcrypt, due to ", err.Error(), "userpayload:", userPayload)
			rw.WriteHeader(http.StatusBadRequest)
			res := uc.Responses(http.StatusBadRequest, "Incorrect Username or password", c.DataMap{"error": "Incorrect Username or password"})
			uc.Json(rw, res)
			return
		}
		rw.WriteHeader(http.StatusOK)
		res := uc.Responses(http.StatusOK, "successfully got it", c.DataMap{"data": User})
		uc.Json(rw, res)

	}
}
