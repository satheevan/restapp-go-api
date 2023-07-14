package routes

import (
	"github.com/gorilla/mux"
	"github.com/pulsarcoder/Projects/restaurantgo/configs"
	"github.com/pulsarcoder/Projects/restaurantgo/controllers"
)

func UserRoute(app *mux.Router) {
	// Just for Testing
	app.HandleFunc("/", configs.Testrusnning()).Methods("GET")
	app.HandleFunc("/restaurant", (&controllers.RestaurantController{}).Index()).Methods("GET")
	// SignUp ***/
	// HandleFunc registers a new route with a matcher for the URL path. See Route.Path() and Route.HandlerFunc().@@
	// Create
	app.HandleFunc("/user", (&controllers.UserController{}).UserCreate()).Methods("POST")

	app.HandleFunc("/login/getuser/", (&controllers.UserController{}).UserGetdata()).Methods("POST")
}
