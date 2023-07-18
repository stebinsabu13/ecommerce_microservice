package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/api/handler"
)

func RegisterAdminRoutes(api *gin.RouterGroup, authHandler handler.AuthHandler, prodHandler handler.ProductHandler) {
	login := api.Group("/admin")
	{
		login.POST("/login", authHandler.AdminLogin)
	}
	signup := api.Group("/admin")
	{
		signup.POST("/signup", authHandler.AdminSignup)
	}
	home := api.Group("/admin")
	{
		home.Use(authHandler.AuthorizationMiddleware("admin"))
		home.POST("/logout", authHandler.AdminLogout)
		product := home.Group("/product")
		{
			product.POST("/add", prodHandler.AddProduct)
			product.GET("/", prodHandler.FindAllProducts)
			productdetail := product.Group("/detail")
			{
				productdetail.GET("/:productid", prodHandler.FindDetailsProductById)
				productdetail.POST("/add", prodHandler.AddProductDetail)
			}
		}
	}
}
