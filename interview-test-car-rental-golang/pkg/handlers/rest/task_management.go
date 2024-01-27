package rest

import (
	"interview-test/car-rental/pkg/handlers"
	"interview-test/car-rental/pkg/models"
	"interview-test/car-rental/pkg/services"
	"interview-test/car-rental/pkg/utilities/standard"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CarRentalHandler struct {
	carRentalSrv services.CarRentalService
}

func NewCarRentalHandler(carRentalSrv services.CarRentalService) handlers.CarRentalAPI {
	return CarRentalHandler{carRentalSrv: carRentalSrv}
}

func (tm CarRentalHandler) CreateCar(c echo.Context) error {
	var bodyRequest models.Car
	if err := c.Bind(&bodyRequest); err != nil {
		return c.JSONPretty(http.StatusBadRequest, standard.BadRequest(err.Error()), "")
	}
	carList, err := tm.carRentalSrv.CreateCar(bodyRequest)
	if err != nil {
		return c.JSONPretty(http.StatusNotFound, standard.BadRequest(err.Error()), "")
	}
	return c.JSONPretty(http.StatusCreated, models.CarResponse{
		StandardReponse: standard.CreateOrUpdateSucccess(),
		Data:            carList,
	}, "")
}

func (tm CarRentalHandler) UpdateCar(c echo.Context) error {
	var bodyRequest models.Car
	if err := c.Bind(&bodyRequest); err != nil {
		return c.JSONPretty(http.StatusBadRequest, standard.BadRequest(err.Error()), "")
	}
	carList, err := tm.carRentalSrv.UpdateCar(bodyRequest)
	if err != nil {
		return c.JSONPretty(http.StatusNotFound, standard.BadRequest(err.Error()), "")
	}
	return c.JSONPretty(http.StatusCreated, models.CarResponse{
		StandardReponse: standard.CreateOrUpdateSucccess(),
		Data:            carList,
	}, "")
}

func (tm CarRentalHandler) DeleteCar(c echo.Context) error {
	carIDQ := c.QueryParam("car_id")
	if carIDQ == "" {
		return c.JSONPretty(http.StatusNotFound, standard.BadRequest("car_id is require"), "")
	}

	carList, err := tm.carRentalSrv.DeleteCar(carIDQ)
	if err != nil {
		return c.JSONPretty(http.StatusNotFound, standard.NotFound(err.Error()), "")
	}
	return c.JSONPretty(http.StatusCreated, models.CarResponse{
		StandardReponse: standard.OkStatus(),
		Data:            carList,
	}, "")
}

func (tm CarRentalHandler) GetCar(c echo.Context) error {
	carIDQ := c.QueryParam("car_id")
	if carIDQ == "" {
		return c.JSONPretty(http.StatusNotFound, standard.BadRequest("car_id is require"), "")
	}

	task, err := tm.carRentalSrv.GetCar(carIDQ)
	if err != nil {
		return c.JSONPretty(http.StatusNotFound, standard.NotFound(err.Error()), "")
	}

	return c.JSONPretty(http.StatusCreated, models.CarResponse{
		StandardReponse: standard.OkStatus(),
		Data:            task,
	}, "")
}

func (tm CarRentalHandler) GetCarList(c echo.Context) error {
	carList, err := tm.carRentalSrv.GetCarList()
	if err != nil {
		return c.JSONPretty(http.StatusNotFound, standard.InternalServerError(err.Error()), "")
	}
	return c.JSONPretty(http.StatusCreated, models.CarResponse{
		StandardReponse: standard.OkStatus(),
		Data:            carList,
	}, "")
}
