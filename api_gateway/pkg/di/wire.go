//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/api"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/api/handler"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/client"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/service"
)

func InitializeAPI(c *config.Config) (*api.Server, error) {
	wire.Build(service.InitClient,
		client.NewauthClient, client.NewProductClient, client.NewCartClient, client.NewOrderClient,
		handler.NewUserHandler, handler.NewProductHandler, handler.NewCartHandler, handler.NewOrderHandler,
		api.NewServerHTTP)
	return &api.Server{}, nil
}
