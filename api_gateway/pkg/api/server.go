package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/api/handler"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/api/routes"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/config"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServerHTTP(c *config.Config, authHandler handler.AuthHandler, productHandler handler.ProductHandler,
	cartHandler handler.CartHandler, orderHandler handler.OrderHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())
	routes.RegisterUserRoutes(engine.Group("/"), authHandler, productHandler, cartHandler, orderHandler)
	routes.RegisterAdminRoutes(engine.Group("/"), authHandler, productHandler)
	return &Server{
		Engine: engine,
		Port:   c.Port,
	}, nil
}

func (c *Server) Start() {
	c.Engine.LoadHTMLGlob("static/*.html")
	c.Engine.Run(c.Port)
}
