package services

import (
	"fmt"
	"interview-test/car-rental/config"
	"interview-test/car-rental/pkg/models"
	"interview-test/car-rental/pkg/utilities"
	"sync"
)

type CarRental struct {
	mu  *sync.Mutex
	cfg config.AppConfig
}

func NewCarRentalService(mutex *sync.Mutex, cfg config.AppConfig) CarRentalService {
	return CarRental{mu: mutex, cfg: cfg}
}

func (cr CarRental) CreateCar(car models.Car) ([]models.Car, error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	carList, _ := utilities.ReadJSONFile(cr.cfg.Config.Filename)

	carIndex := findIndexByName(carList, car.Name)
	if carIndex != -1 {
		return nil, fmt.Errorf("Duplicate car name")
	}

	car.Id = utilities.UUIDGeneratorInstance.GenerateRandomUUID()
	carList = append(carList, car)

	if err := utilities.WriteJSONFile(cr.cfg.Config.Filename, carList); err != nil {
		return nil, err
	}

	return carList, nil
}

func (cr CarRental) UpdateCar(car models.Car) ([]models.Car, error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	carList, err := utilities.ReadJSONFile(cr.cfg.Config.Filename)
	if err != nil {
		return nil, err
	}

	carIndex := findIndexByID(carList, car.Id)
	if carIndex == -1 {
		return nil, fmt.Errorf("Car not found")
	}

	checkName := findIndexByName(carList, car.Name)
	if checkName != -1 && carList[checkName].Id != car.Id {
		return nil, fmt.Errorf("Duplicate car name")
	}

	carList[carIndex] = car

	if err := utilities.WriteJSONFile(cr.cfg.Config.Filename, carList); err != nil {
		return nil, err
	}

	return carList, nil
}

func (cr CarRental) DeleteCar(carID string) ([]models.Car, error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	carList, err := utilities.ReadJSONFile(cr.cfg.Config.Filename)
	if err != nil {
		return nil, err
	}

	carIndex := findIndexByID(carList, carID)
	if carIndex == -1 {
		return nil, fmt.Errorf("Car not found")
	}

	output := deleteElement(carList, carIndex)

	if err := utilities.WriteJSONFile(cr.cfg.Config.Filename, output); err != nil {
		return nil, err
	}

	return output, nil
}

func (cr CarRental) GetCar(carID string) (models.Car, error) {

	carList, err := utilities.ReadJSONFile(cr.cfg.Config.Filename)
	if err != nil {
		return models.Car{}, err
	}

	carIndex := findIndexByID(carList, carID)
	if carIndex == -1 {
		return models.Car{}, fmt.Errorf("Car not found")
	}

	return carList[carIndex], nil
}

func (cr CarRental) GetCarList() ([]models.Car, error) {

	carList, err := utilities.ReadJSONFile(cr.cfg.Config.Filename)
	if err != nil {
		return nil, err
	}

	return carList, nil
}

func findIndexByName(arr []models.Car, target string) int {
	for i, value := range arr {
		if value.Name == target {
			return i
		}
	}
	return -1
}

func findIndexByID(arr []models.Car, target string) int {
	for i, value := range arr {
		if value.Id == target {
			return i
		}
	}
	return -1
}

func deleteElement(arr []models.Car, index int) []models.Car {
	if index < 0 || index >= len(arr) {
		fmt.Println("Index out of bounds")
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}
