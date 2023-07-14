package controllers

import "net/http"

type RestaurantController struct {
	BaseController
}

func (rc RestaurantController) Index() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rc.Json(rw, &struct{ Data string }{
			Data: "restaurants",
		})
	}
}
