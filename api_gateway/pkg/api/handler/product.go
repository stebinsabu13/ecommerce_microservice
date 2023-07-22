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

// LIST PRODUCTS
//
//	@Summary		API FOR LISTING ALL PRODUCTS
//	@Description	LISTING ALL PRODUCTS FROM ADMINS AND USERS END
//	@Tags			ADMIN USER
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int	false	"Enter the page number to display"
//	@Param			limit	query		int	false	"Number of items to retrieve per page"
//	@Success		200		{object}	utils.Response
//	@Failure		401		{object}	utils.Response
//	@Failure		400		{object}	utils.Response
//	@Failure		500		{object}	utils.Response
//	@Router			/user/products [get]
//	@Router			/admin/product [get]
func (cr *ProductHandler) FindAllProducts(c *gin.Context) {
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

// ADD PRODUCT
//
//	@Summary		API FOR ADDING PRODUCT
//	@ID				ADMIN-ADD-PRODUCT
//	@Description	ADDING PRODUCT FROM ADMINS END
//	@Tags			ADMIN
//	@Accept			json
//	@Produce		json
//	@Param			product_details	body		utils.AddProduct	false	"Enter the product details"
//	@Success		200				{object}	utils.Response
//	@Failure		401				{object}	utils.Response
//	@Failure		400				{object}	utils.Response
//	@Failure		500				{object}	utils.Response
//	@Router			/admin/product/add [post]
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

// LIST PRODUCTS DETAILS
//
//	@Summary		API FOR LISTING PRODUCTS DETAILS BY ID
//	@Description	LISTING ALL PRODUCTS DETAILS FROM ADMINS AND USERS END
//	@Tags			ADMIN USER
//	@Accept			json
//	@Produce		json
//	@Param			productid	path		string	true	"Enter the product id that you would like to see the details of"
//	@Param			page		query		int		false	"Enter the page number to display"
//	@Param			limit		query		int		false	"Number of items to retrieve"
//	@Success		200			{object}	utils.Response
//	@Failure		401			{object}	utils.Response
//	@Failure		400			{object}	utils.Response
//	@Failure		500			{object}	utils.Response
//	@Router			/user/products/{productid} [get]
//	@Router			/admin/product/detail/{productid} [get]
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

// ADD PRODUCT DETAILS
//
//	@Summary		API FOR ADDING PRODUCT DETAILS
//	@ID				ADMIN-ADD-PRODUCT-DETAILS
//	@Description	ADDING PRODUCT DETAILS FROM ADMINS END
//	@Tags			ADMIN
//	@Accept			json
//	@Produce		json
//	@Param			product_details	body		utils.AddProductDetail	false	"Enter the product details"
//	@Success		200				{object}	utils.Response
//	@Failure		401				{object}	utils.Response
//	@Failure		400				{object}	utils.Response
//	@Failure		500				{object}	utils.Response
//	@Router			/admin/product/detail/add [post]
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
