package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	// Folder
	// "github.com/pulsarcoder/Projects/restaurantgo/configs"
	"github.com/pulsarcoder/Projects/restaurantgo/routes"
)

func main() {
	router := mux.NewRouter()

	routes.UserRoute(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	fmt.Println("server is connected successfully")
	// configs.ConnectDB()
	log.Fatal(http.ListenAndServe(":6001", handler))

	// log.Fatal(v ...any)
}
