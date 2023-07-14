package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pulsarcoder/Projects/restaurantgo/responses"
	//folders
)

type BaseController struct {
}

func (base BaseController) Json(rw http.ResponseWriter, res interface{}) {
	json.NewEncoder(rw).Encode(res)
}

func (base BaseController) Status400(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusBadRequest)
}
func (base BaseController) Status500(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
}

func (base BaseController) Responses(status int, mgs string, data map[string]interface{}) responses.UserResponse {
	var result responses.UserResponse
	result = responses.UserResponse{
		Status:  status,
		Message: mgs,
		Data:    data,
	}
	return result
}
