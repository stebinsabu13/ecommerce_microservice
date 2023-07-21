package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/api/handler"
)

func RegisterUserRoutes(api *gin.RouterGroup, authHanlder handler.AuthHandler, prodHandler handler.ProductHandler,
	cartHandler handler.CartHandler, orderHandler handler.OrderHandler) {
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
		profile := home.Group("/profile")
		{
			// profile.GET("/", userHandler.ShowUserDetails)
			// profile.PATCH("/edit/details", userHandler.UpdateProfile)
			address := profile.Group("/address")
			{
				address.GET("/", authHanlder.ShowAllAddress)
				address.POST("/add", authHanlder.AddAddress)
			}
		}
		orders := home.Group("/orders")
		{
			orders.GET("/", orderHandler.ShowOrders)
			orders.GET("/detail", orderHandler.ShowOrderDetail)
			// orders.PATCH("/cancel", orderHandler.CancelOrder)
			// orders.PATCH("/return", orderHandler.ReturnOrder)
		}
		cart := home.Group("/cart")
		{
			cart.GET("/", cartHandler.ViewCart)
			cart.PUT("/add", cartHandler.AddtoCart)
			// cart.PUT("/remove", cartHandler.RemovefromCart)
		}
		checkout := home.Group("/checkout")
		{
			checkout.GET("/add", orderHandler.AddtoOrders)
			// checkout.GET("/success", orderHandler.RazorpaymentSuccess)
		}
		// wallet := home.Group("/wallet")
		// {
		// 	wallet.GET("/", userHandler.ViewWallet)
		// }
	}
}
