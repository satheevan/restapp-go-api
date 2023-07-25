package routes

import (
	"github.com/gorilla/mux"
	"github.com/pulsarcoder/Projects/restaurantgo/configs"
	"github.com/pulsarcoder/Projects/restaurantgo/controllers"
	"github.com/pulsarcoder/Projects/restaurantgo/controllers/RestaurantController"
)

// User Routing
func UserRoute(app *mux.Router) {
	// Just for Testing
	app.HandleFunc("/", configs.Testrusnning()).Methods("GET")
	app.HandleFunc("/restaurant", (&RestaurantController.RestaurantController{}).Index()).Methods("GET")
	// SignUp ***/
	// HandleFunc registers a new route with a matcher for the URL path. See Route.Path() and Route.HandlerFunc().@@
	// Create
	app.HandleFunc("/user-create", (&controllers.UserController{}).UserCreate()).Methods("POST")
	app.HandleFunc("/login/getuser/{email}", (&controllers.UserController{}).UserGetdata()).Methods("GET")

	//Login Route
	app.HandleFunc("/login", (&controllers.UserController{}).LoginUser()).Methods("POST")
}

// Restaurant Routing
func RestaurantRoute(app *mux.Router) {
	app.HandleFunc("/restaurant-create", (&RestaurantController.RestaurantController{}).CreateRestaurants()).Methods("Post")

	app.HandleFunc("/restaurant-list", (&RestaurantController.RestaurantController{}).GetAllRestaurants()).Methods("Get")
	app.HandleFunc("/restaurant-getonedata/{id}", (&RestaurantController.RestaurantController{}).UpdateOneRestaurants()).Methods("Get")
}
