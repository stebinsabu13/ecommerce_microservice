package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/api/handler"
)

func RegisterUserRoutes(api *gin.RouterGroup, authHanlder handler.AuthHandler, prodHandler handler.ProductHandler) {
	signup := api.Group("/user")
	{
		signup.POST("/signup", authHanlder.Signup)
		signup.POST("/signup/otp/verify", authHanlder.SignUpOtpVerify)
	}
	Login := api.Group("/user")
	{
		Login.POST("/login", authHanlder.UserLogin)
	}
	home := api.Group("/user")
	{
		home.Use(authHanlder.AuthorizationMiddleware("user"))
		home.POST("/logout", authHanlder.UserLogout)
		product := home.Group("/products")
		{
			product.GET("/", prodHandler.FindAllProducts)
			product.GET("/:productid", prodHandler.FindDetailsProductById)
		}
	}
}
