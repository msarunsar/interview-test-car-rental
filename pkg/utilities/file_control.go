package utilities

import (
	"encoding/json"
	"interview-test/car-rental/pkg/models"
	"os"
)

func ReadJSONFile(filename string) ([]models.Car, error) {
	var config []models.Car

	// Read the JSON file
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	// Unmarshal JSON data into the Config struct
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func WriteJSONFile(filename string, cars []models.Car) error {
	// Marshal the Config struct into JSON
	jsonData, err := json.MarshalIndent(cars, "", "    ")
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = os.WriteFile(filename, jsonData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
