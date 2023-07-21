//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/api"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/api/service"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/client"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/db"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/repository"
)

func InitializeServe(c *config.Config) (*api.Server, error) {
	wire.Build(db.Initdb,
		client.InitClient, client.NewProdClient,
		repository.NewCartRepo,
		service.NewCartService,
		api.NewgrpcServe)
	return &api.Server{}, nil
}
