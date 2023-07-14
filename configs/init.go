package configs

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Newclient func() *mongo.Client

type Database struct {
}

// the func from Routes
func Testrusnning() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		json.NewEncoder(rw).Encode(map[string]string{
			"data": "Hello the server is running successfully",
		})
	}
}
