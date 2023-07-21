// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/api"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/api/service"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/client"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/db"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/repository"
)

// Injectors from wire.go:

func InitializeServe(c *config.Config) (*api.Server, error) {
	gormDB, err := db.Initdb(c)
	if err != nil {
		return nil, err
	}
	cartRepo := repository.NewCartRepo(gormDB)
	productServiceClient, err := client.InitClient(c)
	if err != nil {
		return nil, err
	}
	productClient := client.NewProdClient(productServiceClient)
	cartServiceServer := service.NewCartService(cartRepo, productClient)
	server, err := api.NewgrpcServe(c, cartServiceServer)
	if err != nil {
		return nil, err
	}
	return server, nil
}
