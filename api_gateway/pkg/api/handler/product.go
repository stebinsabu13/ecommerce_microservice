package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	client "github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/utils"
)

type ProductHandler struct {
	Client client.ProductClient
}

func NewProductHandler(client client.ProductClient) ProductHandler {
	return ProductHandler{
		Client: client,
	}
}

func (cr *ProductHandler) FindAllProducts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, err1 := strconv.Atoi(c.DefaultQuery("limit", "1"))
	err = errors.Join(err, err1)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	offset := (page - 1) * limit
	pagination := utils.Pagination{
		Offset: uint(offset),
		Limit:  uint(limit),
	}
	res, err := cr.Client.FindAllProducts(c.Request.Context(), pagination)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"Products_list": res.Product,
	})
}

func (cr *ProductHandler) AddProduct(c *gin.Context) {
	var body utils.AddProduct
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := cr.Client.AddProduct(c.Request.Context(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"Success": res.Reponse,
	})
}

func (cr *ProductHandler) FindDetailsProductById(c *gin.Context) {
	id := c.Param("productid")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, err1 := strconv.Atoi(c.DefaultQuery("limit", "5"))
	err = errors.Join(err, err1)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	offset := (page - 1) * limit
	pagination := utils.Pagination{
		Offset: uint(offset),
		Limit:  uint(limit),
	}
	res, err := cr.Client.FindProductById(c.Request.Context(), id, pagination)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"Details": res.ProductDetail,
	})
}

func (cr *ProductHandler) AddProductDetail(c *gin.Context) {
	var body utils.AddProductDetail
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := cr.Client.AddProductDetail(c.Request.Context(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"Success": res.Reponse,
	})
}
