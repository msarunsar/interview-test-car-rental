package main

import (
	"fmt"
	"interview-test/car-rental/config"
	"interview-test/car-rental/pkg/handlers/rest"

	"github.com/labstack/gommon/log"
)

func main() {

	var cfg config.AppConfig
	if err := config.InitCFG(&cfg); err != nil {
		log.Fatal("Fatal error: ", err)
	}

	e := rest.InitRouter(cfg)
	address := fmt.Sprintf("localhost:%d", cfg.App.Port)
	e.Logger.Fatal(e.Start(address))
}
