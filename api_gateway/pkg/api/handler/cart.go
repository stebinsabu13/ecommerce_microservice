package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/client/interfaces"
)

type CartHandler struct {
	Client interfaces.CartClient
}

func NewCartHandler(client interfaces.CartClient) CartHandler {
	return CartHandler{
		Client: client,
	}
}

func (cr *CartHandler) ViewCart(c *gin.Context) {
	id, ok := c.Get("user-id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not ok",
		})
		return
	}
	res, err := cr.Client.ViewCart(id.(uint))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(res.Grandtotal)
	c.JSON(int(res.Status), gin.H{
		"Success": res,
	})
}

func (cr *CartHandler) AddtoCart(c *gin.Context) {
	prodetid := c.Query("id")
	id, ok := c.Get("user-id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not ok",
		})
		return
	}
	res, err := cr.Client.AdditemtoCart(prodetid, id.(uint))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"Success": res.Response,
	})
}
