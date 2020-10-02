package data

import (
	"fmt"
	"time"
)

// Plant struct represent the plant  details
// swagger:model
type Plant struct {
	// the id for this plant
	//
	// required: true
	// min: 1
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float32 `json:"price" validate:"gt=0"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

// Initialize the plants list
type Plants []*Plant

// Initialize Excceptions
var PlantNotFoundException = fmt.Errorf("Plant not found")

// GetPlants is used to get the plants data list
func GetPlants() Plants {
	return plantsList
}

// GetPlant is used to get the plant data based on plant id
func GetPlant(id int) (*Plant, error) {
	plant, _, notFoundError := getPlantPosition(id)
	return plant, notFoundError
}

// AddPlant is used to insert the newplant data into plants data list
func AddPlant(plant *Plant) {
	plant.ID = plantsList[len(plantsList)-1].ID + 1
	plantsList = append(plantsList, plant)
}

// UpdatePlant is used to update the plant data into plants data list based on plant id
func UpdatePlant(id int, plant *Plant) error {
	updatePlant, position, notFoundError := getPlantPosition(id)

	if notFoundError != nil {
		return PlantNotFoundException
	}

	plant.ID = updatePlant.ID
	plantsList[position] = plant
	return nil
}

// DeletePlant is used to delete the plant data into plants data list based on plant id
func DeletePlant(id int) error {
	_, position, notFoundError := getPlantPosition(id)

	if notFoundError != nil {
		return PlantNotFoundException
	}
	plantsList = append(plantsList[:position], plantsList[position+1:]...)

	return nil
}

// getPlantPosition is used to get the plant data and position from plants data list
func getPlantPosition(id int) (*Plant, int, error) {
	for position, plantData := range plantsList {
		if plantData.ID == id {
			return plantData, position, nil
		}
	}
	return nil, -1, PlantNotFoundException
}

// Create temporary plant data
var plantsList = []*Plant{
	&Plant{
		ID:          1,
		Name:        "Rose",
		Description: "Beautiful Flower",
		Category:    "Flower",
		Price:       100.00,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   time.Now().UTC().String(),
	},
	&Plant{
		ID:          2,
		Name:        "Apple",
		Description: "Tasty Fruit",
		Category:    "Fruit",
		Price:       500.00,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   time.Now().UTC().String(),
	},
}
