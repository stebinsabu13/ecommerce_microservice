package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	client "github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/utils"
)

type AuthHandler struct {
	Client client.AuthClient
}

func NewUserHandler(client client.AuthClient) AuthHandler {
	return AuthHandler{
		Client: client,
	}
}

func (cr *AuthHandler) Signup(c *gin.Context) {
	var body utils.BodySignUpuser
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "binding error" + err.Error(),
		})
		return
	}
	res, err := cr.Client.UserSignup(context.Background(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"responseid":  &res.Responseid,
		"referalcode": body.ReferalCode,
	})
}

func (cr *AuthHandler) SignUpOtpVerify(c *gin.Context) {
	var body utils.Otpverify
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error binding json",
		})
		return
	}
	res, err := cr.Client.UserSignupOtpVerify(context.Background(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), &res)
}

func (cr *AuthHandler) UserLogin(c *gin.Context) {
	var body utils.BodyLogin
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error binding json",
		})
		return
	}
	res, err := cr.Client.UserLogin(context.Background(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}
	c.SetCookie("user-token", res.Token, int(time.Now().Add(1*time.Hour).Unix()), "/", "localhost", false, true)
	c.JSON(int(res.Status), gin.H{
		"User": res.User,
	})
}

func (cr *AuthHandler) UserLogout(c *gin.Context) {
	c.SetCookie("user-token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"Logout": "Success",
	})
}

func (cr *AuthHandler) AdminSignup(c *gin.Context) {
	var body utils.BodySignUpuser
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "binding error" + err.Error(),
		})
		return
	}
	res, err := cr.Client.AdminSignup(context.Background(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"response": res.Response,
	})
}

func (cr *AuthHandler) AdminLogin(c *gin.Context) {
	var body utils.BodyLogin
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error binding json",
		})
		return
	}
	res, err := cr.Client.AdminLogin(context.Background(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}
	c.SetCookie("admin-token", res.Token, int(time.Now().Add(1*time.Hour).Unix()), "/", "localhost", false, true)
	c.JSON(int(res.Status), gin.H{
		"User": res.User,
	})
}

func (cr *AuthHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("admin-token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"Logout": "Success",
	})
}
