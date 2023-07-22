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

// LIST CART_DETAILS
//
//	@Summary		API FOR DISPLAYING CART TO USER
//	@ID				USER-LIST-CART
//	@Description	LISTING CART AND ITEMS FROM USERS END
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.Response
//	@Failure		401	{object}	utils.Response
//	@Failure		400	{object}	utils.Response
//	@Failure		500	{object}	utils.Response
//	@Router			/user/cart [get]
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

// ADD TO CART
//
//	@Summary		API FOR ADDING PRODUCTS TO CART BY USER
//	@ID				USER-ADD-TO-CART
//	@Description	ADDING ITEMS TO CART FROM USERS END
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"Enter the product id"
//	@Success		200	{object}	utils.Response
//	@Failure		401	{object}	utils.Response
//	@Failure		400	{object}	utils.Response
//	@Failure		500	{object}	utils.Response
//	@Router			/user/cart/add [put]
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
