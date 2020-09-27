package data

import "testing"

func TestPlantStructValidation(testcase *testing.T) {
	plant := &Plant{
		Name:  "apple",
		Price: 200.00,
	}

	validationError := plant.Validate()
	if validationError != nil {
		testcase.Fatal(validationError)
	}
}
