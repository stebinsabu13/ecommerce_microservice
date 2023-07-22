package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/client/interfaces"
)

type OrderHandler struct {
	Client interfaces.OrderClient
}

func NewOrderHandler(client interfaces.OrderClient) OrderHandler {
	return OrderHandler{
		Client: client,
	}
}

// PLACE A NEW ORDER
//
//	@Summary		API FOR PLACING A NEW ORDER
//	@ID				USER-PROCEED-ORDER
//	@Description	Users can place a new order with the cart items.
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			paymentid	query		string	true	"Enter the payment id"
//	@Param			addressid	query		string	true	"Enter the address id"
//	@Param			code		query		string	false	"If you have a coupon,Enter the coupon code"
//	@Success		200			{object}	utils.Response
//	@Failure		401			{object}	utils.Response
//	@Failure		400			{object}	utils.Response
//	@Failure		500			{object}	utils.Response
//	@Router			/user/checkout/add [get]
func (cr *OrderHandler) AddtoOrders(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	addressid, _ := strconv.Atoi(c.Query("addressid"))
	paymentid, _ := strconv.Atoi(c.Query("paymentid"))
	id, ok := c.Get("user-id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not ok",
		})
		return
	}
	res, err := cr.Client.AddtoOrders(c.Request.Context(), uint(addressid), uint(paymentid), id.(uint), code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if paymentid == 2 {
		c.HTML(200, "app.html", gin.H{
			"UserID":      res.UserID,
			"Orderid":     res.RazorpayOrderID,
			"Total_price": res.AmountToPay,
		})
	} else {
		c.JSON(int(res.Status), gin.H{
			"Success": res.Response,
		})
	}
}

// VIEW ORDERS
//
//	@Summary		API FOR VIEWING ORDERS
//	@Description	Users can view all orders.
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.Response
//	@Failure		401	{object}	utils.Response
//	@Failure		400	{object}	utils.Response
//	@Failure		500	{object}	utils.Response
//	@Router			/user/orders [get]
func (cr *OrderHandler) ShowOrders(c *gin.Context) {
	id, ok := c.Get("user-id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not ok",
		})
		return
	}
	res, err := cr.Client.Orders(c.Request.Context(), id.(uint))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"ORDERS": res.Orders,
	})
}

// VIEW ORDERS DETAILS
//
//	@Summary		API FOR VIEWING ORDERS DETAILS
//	@Description	Users can the selected order details.
//	@Tags			ADMIN USER
//	@Accept			json
//	@Produce		json
//	@Param			orderid	query		uint	true	"Enter the order id"
//	@Success		200		{object}	utils.Response
//	@Failure		401		{object}	utils.Response
//	@Failure		400		{object}	utils.Response
//	@Failure		500		{object}	utils.Response
//	@Router			/user/orders/detail [get]
//	@Router			/admin/order/detail [get]
func (cr *OrderHandler) ShowOrderDetail(c *gin.Context) {
	id, err1 := strconv.Atoi(c.Query("orderid"))
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err1.Error(),
		})
		return
	}
	res, err := cr.Client.OrderDetail(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"ORDER DETAILS": res.Details,
	})
}
