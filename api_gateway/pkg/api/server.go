package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/stebinsabu13/ecommerce_microservice/api_gateway/docs"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/api/handler"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/api/routes"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServerHTTP(c *config.Config, authHandler handler.AuthHandler, productHandler handler.ProductHandler,
	cartHandler handler.CartHandler, orderHandler handler.OrderHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.RegisterUserRoutes(engine.Group("/"), authHandler, productHandler, cartHandler, orderHandler)
	routes.RegisterAdminRoutes(engine.Group("/"), authHandler, productHandler)
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"StatusCode": 404,
			"msg":        "invalid url",
		})
	})
	return &Server{
		Engine: engine,
		Port:   c.Port,
	}, nil
}

func (c *Server) Start() {
	c.Engine.LoadHTMLGlob("static/*.html")
	c.Engine.Run(c.Port)
}
