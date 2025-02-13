package main

import (
	"fmt"
	"log"
	"os"
	"test-case-roketin/common/constants"
	routeRegisters "test-case-roketin/routes/route-registers"
	"test-case-roketin/utils/database/postgres"
	"test-case-roketin/utils/env"
	"test-case-roketin/utils/route"
)

func main() {
	//Init config app.env
	config, err := env.LoadConfig(".")
	if err != nil {
		message := fmt.Sprintf("%s can't load configuration file", constants.Configuration)
		log.Fatal(message)
	}

	//Set timezone
	err = os.Setenv("TZ", config.DbTz)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error set timezone: %s", err.Error()))
	}

	//Init database
	db := postgres.ConfigurationPostgres(config)

	//Init router
	router := route.InitRouter(config)

	//Register routes
	apiV1 := router.Group("api/v1")
	routeRegisters.RouteRegister(db, apiV1)

	//Run route
	route.RunRoute(config, router)
}
