package main

import (
	"log"

	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/di"
)

//	@title			E-COMMERCE MICROSERVICE
//	@version		2.0
//	@description	SPORTZONE_E-COMMERCE MICROSERVICE API built using Go, PSQL, REST API following Clean Architecture.

//	@contact
// name: Stebin Sabu
// url: https://github.com/stebinsabu13
// email: stebinsabu369@gmail.com

//	@license
// name: MIT
// url: https://opensource.org/licenses/MIT

//	@host	localhost:4000

// @Basepath	/
// @Accept		json
// @Produce	json
// @Router		/ [get]
func main() {
	c, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("cannot load config:", configerr)
	}
	server, dierr := di.InitializeAPI(c)
	if dierr != nil {
		log.Fatal("cannot initialize server", dierr)
	}
	server.Start()
}
