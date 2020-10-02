package handlers

import (
	"net/http"

	"github.com/saravase/golang_mux_swagger/plant-api/data"
)

// swagger:route POST /plant plants addPlant
// responses:
// 		200: successContent
//		422: errorValidation
//		400: errorResponse
// AddPlant used to insert the new plant data into the datastore
func (plant *Plant) AddPlant(response http.ResponseWriter, request *http.Request) {
	plant.logger.Printf("[DEBUG] Add the plant data")

	plantData := request.Context().Value(KeyPlant{}).(*data.Plant)
	plant.logger.Printf("[DEBUG] Added plant data: %#v\n", plantData)

	data.AddPlant(plantData)

}
