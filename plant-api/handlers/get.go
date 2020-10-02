package handlers

import (
	"net/http"

	"github.com/saravase/golang_mux_swagger/plant-api/data"
)

// swagger:route GET /plant plants listPlants
// Returns a list of plants
// responses:
//   200: plantsResponse

//GetPlants is used to fetch all the plants data from the datastore
func (plant *Plant) GetPlants(response http.ResponseWriter, request *http.Request) {
	plant.logger.Printf("[DEBUG] Get all the plant data")
	plants := data.GetPlants()
	marshalError := data.ToJSON(plants, response)

	if marshalError != nil {
		plant.logger.Printf("[ERROR] Serializing plant data : %s", marshalError)
	}
}

// swagger:route GET /plant/{id} plants singlePlant
// Return plant data based on plant id
// responses:
//   200: plantResponse
//   404: errorResponse

//GetPlant is used to fetch the plant data based on plant id from the datastore
func (plant *Plant) GetPlant(response http.ResponseWriter, request *http.Request) {
	plant.logger.Printf("[DEBUG] Fetch plant data based on plant id")

	id := GetPlantID(request)
	plantData, err := data.GetPlant(id)

	switch err {
	case nil:
	case data.PlantNotFoundException:
		plant.logger.Printf("[ERROR] Fetch plant based on plant id")
		response.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, response)
		return

	default:
		plant.logger.Printf("[ERROR] Fetch plant based on plant id")
		response.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, response)
		return
	}

	marshalError := data.ToJSON(plantData, response)

	if marshalError != nil {
		plant.logger.Printf("[ERROR] Serializing plant data : %s", marshalError)
	}
}
