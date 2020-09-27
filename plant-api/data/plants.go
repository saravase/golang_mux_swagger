package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

// Plant struct represent the plant  details
type Plant struct {
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

// Serialize the raw data format into JSON format
func (plants *Plants) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(plants)
}

// Deserialize the JSON format into raw data format
func (plant *Plant) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(plant)
}

// Validate validate the plant struct
func (plant *Plant) Validate() error {
	//New returns a new instance of 'validate' with sane defaults
	plantValidator := validator.New()

	//Struct validates a structs exposed fields, and automatically validates nested structs
	return plantValidator.Struct(plant)
}

// GetAllPlants is used to get the plants data list
func GetAllPlants() Plants {
	return plantsList
}

// AddPlant is used to insert the newplant data into plants data list
func AddPlant(plant *Plant) {
	plant.ID = generatePlantNextID()
	plantsList = append(plantsList, plant)
}

// generatePlantNextID is used to generate the id for the newly added plant
func generatePlantNextID() int {
	plant := plantsList[len(plantsList)-1]
	return plant.ID + 1
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
