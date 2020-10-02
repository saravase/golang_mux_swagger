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
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/saravase/golang_mux_swagger/plant-api/data"
)

// Initialize struct type Plant with properties
type Plant struct {
	logger     *log.Logger
	validation *data.Validation
}

// GenericError - Generic error return from server
type GenericError struct {
	Message string `json: "message"`
}

// ValidationError - Collection of validation error return from validator
type ValidationError struct {
	Messages []string `json: "messages"`
}

type KeyPlant struct{}

// Initialize the Plant struct properties
func NewPlant(logger *log.Logger, validation *data.Validation) *Plant {
	return &Plant{
		logger:     logger,
		validation: validation,
	}
}

// Get the plant id from the request URL path
func GetPlantID(request *http.Request) int {
	// Parse the id from URL path
	urlInfo := mux.Vars(request)
	id, err := strconv.Atoi(urlInfo["id"])

	if err != nil {
		// This is should not occur
		panic(err)
	}

	return id
}
