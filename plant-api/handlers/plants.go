// Package classification of Plant API
//
// Documentation for Plant API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/saravase/golang_mux_swagger/plant-api/data"

	"github.com/gorilla/mux"
)

// Initialize struct type Plant with properties
type Plant struct {
	logger *log.Logger
}

// Create struct type Plant with properties
func NewPlant(logger *log.Logger) *Plant {
	return &Plant{
		logger,
	}
}

//getPlants is used to fetch all the plants data from the datastore
func (plant *Plant) GetPlants(response http.ResponseWriter, request *http.Request) {
	plantsList := data.GetAllPlants()
	marshalError := plantsList.ToJSON(response)

	if marshalError != nil {
		plant.logger.Printf("While, Marshaling the plant data. Reason : %s", marshalError)
		http.Error(response, "JSON marshaling failed.", http.StatusInternalServerError)
	}
}

//createPlants used to insert the new plant data into the datastore
func (plant *Plant) CreatePlant(response http.ResponseWriter, request *http.Request) {

	plantData := request.Context().Value(KeyPlant{}).(data.Plant)

	data.AddPlant(&plantData)
	plantsList := &data.Plants{&plantData}
	plantsList.ToJSON(response)
}

//updatePlant used to update the plant data into the datastore based on id.
func (plant *Plant) UpdatePlant(response http.ResponseWriter, request *http.Request) {

	requestInfo := mux.Vars(request)
	id, convertionError := strconv.Atoi(requestInfo["id"])

	if convertionError != nil {
		http.Error(response, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	plantData := request.Context().Value(KeyPlant{}).(data.Plant)
	updateError := data.UpdatePlant(id, &plantData)

	if updateError != nil {
		plant.logger.Printf("While, Update the plant data. Reason : %s", updateError)
		http.Error(response, "Plant not found.", http.StatusNotFound)
		return
	}
	response.WriteHeader(http.StatusOK)
}

//deletePlant used to delete the plant data into the datastore based on id.
func (plant *Plant) DeletePlant(response http.ResponseWriter, request *http.Request) {

	requestInfo := mux.Vars(request)
	id, convertionError := strconv.Atoi(requestInfo["id"])

	if convertionError != nil {
		http.Error(response, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	deleteError := data.DeletePlant(id)

	if deleteError != nil {
		plant.logger.Printf("While, Delete the plant data. Reason : %s", deleteError)
		http.Error(response, "Plant not found.", http.StatusNotFound)
		return
	}
	response.WriteHeader(http.StatusOK)
}

type KeyPlant struct{}

func (plant *Plant) PlantValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		plantData := data.Plant{}
		marshalError := plantData.FromJSON(request.Body)

		if marshalError != nil {
			plant.logger.Printf("While, Marshaling the plant data. Reason : %s", marshalError)
			http.Error(response, "JSON Unmarshaling failed.", http.StatusBadRequest)
			return
		}

		// validate the plant
		validationError := plantData.Validate()
		if validationError != nil {
			plant.logger.Println("[ERROR] Validating plant ", validationError)
			http.Error(
				response,
				fmt.Sprintf("Error validating plant. Reason : %s", validationError),
				http.StatusBadRequest,
			)
			return
		}

		// add the plant to the context
		plantContext := context.WithValue(request.Context(), KeyPlant{}, plantData)
		request = request.WithContext(plantContext)

		// call the next handler, which can be another middlware or final handler
		next.ServeHTTP(response, request)
	})
}
