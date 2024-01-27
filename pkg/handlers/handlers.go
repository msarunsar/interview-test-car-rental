package handlers

import "github.com/labstack/echo/v4"

type CarRentalAPI interface {
	CreateCar(c echo.Context) error
	UpdateCar(c echo.Context) error
	DeleteCar(c echo.Context) error
	GetCar(c echo.Context) error
	GetCarList(c echo.Context) error
}
