package handlers

import (
	"net/http"

	"github.com/saravase/golang_mux_swagger/plant-api/data"
)

// swagger:route DELETE /plant/{id} plants deletePlant
// Returns empty content
// responses:
//   204: noContent
//   404: errorResponse

//DeletePlant used to delete the plant data into the datastore based on id.
func (plant *Plant) DeletePlant(response http.ResponseWriter, request *http.Request) {

	plant.logger.Printf("[DEBUG] Delete plant data based on plant id")

	id := GetPlantID(request)
	err := data.DeletePlant(id)

	switch err {
	case nil:
	case data.PlantNotFoundException:
		plant.logger.Printf("[ERROR] Delete plant based on plant id")
		response.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, response)
		return

	default:
		plant.logger.Printf("[ERROR] Delete plant based on plant id")
		response.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, response)
		return
	}

	response.WriteHeader(http.StatusNoContent)

}
