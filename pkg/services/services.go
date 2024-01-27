package services

import "interview-test/car-rental/pkg/models"

type CarRentalService interface {
	CreateCar(car models.Car) ([]models.Car, error)
	UpdateCar(car models.Car) ([]models.Car, error)
	DeleteCar(carID string) ([]models.Car, error)
	GetCar(carID string) (models.Car, error)
	GetCarList() ([]models.Car, error)
}
