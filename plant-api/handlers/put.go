package handlers

import (
	"net/http"

	"github.com/saravase/golang_mux_swagger/plant-api/data"
)

// swagger:route PUT /plant/{id} plants updatePlant
// responses:
// 		200: successContent
//		422: errorValidation
//		400: errorResponse
// UpdatePlant used to update the plant data into the datastore based on id.
func (plant *Plant) UpdatePlant(response http.ResponseWriter, request *http.Request) {

	plant.logger.Printf("[DEBUG] Update the plant data")

	id := GetPlantID(request)
	plantData := request.Context().Value(KeyPlant{}).(*data.Plant)
	err := data.UpdatePlant(id, plantData)

	switch err {
	case nil:
	case data.PlantNotFoundException:
		plant.logger.Printf("[ERROR] Update plant based on plant id")
		response.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, response)
		return

	default:
		plant.logger.Printf("[ERROR] Update plant based on plant id")
		response.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, response)
		return
	}

	plant.logger.Printf("[DEBUG] Updated plant data: %#v\n", plantData)
	response.WriteHeader(http.StatusOK)
}
