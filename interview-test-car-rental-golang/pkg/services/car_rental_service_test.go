package services

import (
	"interview-test/car-rental/config"
	"interview-test/car-rental/pkg/models"
	"interview-test/car-rental/pkg/utilities"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CarRentalTestSuit struct {
	suite.Suite
	underTest CarRentalService
}

func (s *CarRentalTestSuit) SetupTest() {
	var mutex sync.Mutex
	var cfg config.AppConfig
	cfg.Config.Filename = "cars-test.json"
	s.underTest = NewCarRentalService(&mutex, cfg)
	utilities.WriteJSONFile(cfg.Config.Filename, []models.Car{})
}

func TestCarRentalTestSuite(t *testing.T) {
	suite.Run(t, new(CarRentalTestSuit))
}

type MockUUIDGenerator struct{}

func (m MockUUIDGenerator) GenerateRandomUUID() string {
	return "mocked-uuid"
}

func (s *CarRentalTestSuit) TestCarRental() {
	originalGenerator := utilities.UUIDGeneratorInstance
	defer func() {
		utilities.UUIDGeneratorInstance = originalGenerator
	}()

	utilities.UUIDGeneratorInstance = MockUUIDGenerator{}

	testData := models.Car{
		Name:     "testcar",
		Price:    200000,
		Discount: 10,
	}

	var testCreate = struct {
		mock   models.Car
		expect models.Car
	}{
		mock: testData,
		expect: models.Car{
			Id:       "mocked-uuid",
			Name:     "testcar",
			Price:    200000,
			Discount: 10,
		},
	}

	s.Run("create", func() {
		result, err := s.underTest.CreateCar(testCreate.mock)
		s.NoError(err)
		s.Equal(testCreate.expect, result[0])
	})

	testData.Id = "mocked-uuid"
	testData.Name = "testcar2"

	var testUpdate = struct {
		mock   models.Car
		expect models.Car
	}{
		mock: testData,
		expect: models.Car{
			Id:       "mocked-uuid",
			Name:     "testcar2",
			Price:    200000,
			Discount: 10,
		},
	}

	s.Run("update", func() {
		result, err := s.underTest.UpdateCar(testUpdate.mock)
		s.NoError(err)
		s.Equal(testUpdate.expect, result[0])
	})

	s.Run("get", func() {
		result, err := s.underTest.GetCar(testData.Id)
		s.NoError(err)
		s.Equal(testUpdate.expect, result)
	})

	s.Run("list", func() {
		result, err := s.underTest.GetCarList()
		s.NoError(err)
		s.Equal(testUpdate.expect, result[0])
	})

	s.Run("delete", func() {
		result, err := s.underTest.DeleteCar(testData.Id)
		s.NoError(err)
		s.Equal(0, len(result))
	})

}
