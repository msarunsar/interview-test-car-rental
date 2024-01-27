package rest

import (
	"interview-test/car-rental/config"
	"interview-test/car-rental/pkg/services"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(cfg config.AppConfig) *echo.Echo {
	e := echo.New()

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Check if the provided username and password are valid
		if username == cfg.App.Authentication.Username && password == cfg.App.Authentication.Password {
			return true, nil
		}
		return false, nil
	}))

	main := e.Group("/car-rental")
	api := main.Group("/api")
	g := api.Group("/v1")

	var mutex sync.Mutex
	carRentalSrv := services.NewCarRentalService(&mutex, cfg)
	initTaskManagementRouter(g, carRentalSrv)
	return e
}

func initTaskManagementRouter(e *echo.Group, carRentalSrv services.CarRentalService) {
	handler := NewCarRentalHandler(carRentalSrv)

	e.POST("/car", handler.CreateCar)

	e.PUT("/car", handler.UpdateCar)

	e.DELETE("/car", handler.DeleteCar)

	e.GET("/car", handler.GetCar)

	e.GET("/cars", handler.GetCarList)
}
