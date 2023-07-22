package handler

import (
	"context"
	"fmt"
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

// USER SIGN-UP WITH SENDING OTP
//
//	@Summary		API FOR NEW USER SIGN UP
//	@ID				SIGNUP-USER
//	@Description	CREATE A NEW USER WITH REQUIRED DETAILS
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			user_details	body		utils.BodySignUpuser	false	"New user Details"
//	@Success		200				{object}	utils.Response
//	@Failure		400				{object}	utils.Response
//	@Failure		500				{object}	utils.Response
//	@Router			/user/signup [post]
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

// USER SIGN-UP WITH VERIFICATION OF OTP
//
//	@Summary		API FOR NEW USER SIGN UP OTP VERIFICATION
//	@ID				SIGNUP-USER-OTP-VERIFY
//	@Description	VERIFY THE OTP AND UPDATE THE VERIFIED COLUMN
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			otp_details	body		utils.Otpverify	false	"otp"
//	@Success		200			{object}	utils.Response
//	@Failure		401			{object}	utils.Response
//	@Failure		400			{object}	utils.Response
//	@Failure		500			{object}	utils.Response
//	@Router			/user/signup/otp/verify [post]
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

// USER USERLOGIN
//
//	@Summary		API FOR USER LOGIN
//	@ID				USER-LOGIN
//	@Description	VERIFY THE EMAIL,PASSWORD AND GENERATE A JWT TOKEN AND SET IT TO A COOKIE
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			login_details	body		utils.BodyLogin	true	"Enter the email and password"
//	@Success		200				{object}	utils.Response
//	@Failure		401				{object}	utils.Response
//	@Failure		400				{object}	utils.Response
//	@Failure		500				{object}	utils.Response
//	@Router			/user/login [post]
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

// USERLOGOUT
//
//	@Summary		API FOR USER LOGOUT
//	@ID				USER-LOGOUT
//	@Description	LOGOUT USER AND ALSO CLEAR COOKIES
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.Response
//	@Failure		401	{object}	utils.Response
//	@Failure		400	{object}	utils.Response
//	@Failure		500	{object}	utils.Response
//	@Router			/user/logout [post]
func (cr *AuthHandler) UserLogout(c *gin.Context) {
	c.SetCookie("user-token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"Logout": "Success",
	})
}

// ADMIN SIGN-UP
//
//	@Summary		API FOR NEW ADMIN SIGN UP
//	@ID				SIGNUP-ADMIN
//	@Description	CREATE A NEW ADMIN WITH REQUIRED DETAILS
//	@Tags			ADMIN
//	@Accept			json
//	@Produce		json
//	@Param			admin_details	body		utils.BodySignUpuser	false	"New Admin Details"
//	@Success		200				{object}	utils.Response
//	@Failure		400				{object}	utils.Response
//	@Failure		400				{object}	utils.Response
//	@Failure		500				{object}	utils.Response
//	@Router			/admin/signup [post]
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

// ADMIN LOGIN
//
//	@Summary		API FOR ADMIN LOGIN
//	@ID				ADMIN-LOGIN
//	@Description	VERIFY THE EMAIL,PASSWORD AND GENERATE A JWT TOKEN AND SET IT TO A COOKIE
//	@Tags			ADMIN
//	@Accept			json
//	@Produce		json
//	@Param			login_details	body		utils.BodyLogin	true	"Enter the email and password"
//	@Success		200				{object}	utils.Response
//	@Failure		400				{object}	utils.Response
//	@Failure		401				{object}	utils.Response
//	@Failure		500				{object}	utils.Response
//	@Router			/admin/login [post]
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

// ADMIN LOGOUT
//
//	@Summary		API FOR ADMIN LOGOUT
//	@ID				ADMIN-LOGOUT
//	@Description	ADMIN LOGOUT
//	@Tags			ADMIN
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.Response
//	@Failure		401	{object}	utils.Response
//	@Failure		400	{object}	utils.Response
//	@Failure		500	{object}	utils.Response
//	@Router			/admin/logout [post]
func (cr *AuthHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("admin-token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"Logout": "Success",
	})
}

// LIST ADDRESS
//
//	@Summary		API FOR LISTING ADDRESSES
//	@ID				USER-LIST-ADDRESS
//	@Description	LISTING ALL ADDRESSES FOR THE PARTICULAR USER
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.Response
//	@Failure		401	{object}	utils.Response
//	@Failure		400	{object}	utils.Response
//	@Failure		500	{object}	utils.Response
//	@Router			/user/profile/address [get]
func (cr *AuthHandler) ShowAllAddress(c *gin.Context) {
	id, ok := c.Get("user-id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not ok",
		})
		return
	}
	addresses, err := cr.Client.ShowAddress(c.Request.Context(), id.(uint))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"addresses": addresses,
	})
}

// ADD ADDRESS
//
//	@Summary		API FOR ADDING ADDRESS
//	@ID				USER-ADD-ADDRESS
//	@Description	ADDING NEW ADDRESS TO USER PROFILE
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			address_details	body		utils.AddAddress	true	"Add the address details"
//	@Success		200				{object}	utils.Response
//	@Failure		401				{object}	utils.Response
//	@Failure		400				{object}	utils.Response
//	@Failure		500				{object}	utils.Response
//	@Router			/user/profile/address/add [post]
func (cr *AuthHandler) AddAddress(c *gin.Context) {
	var body utils.AddAddress
	id, ok := c.Get("user-id")
	fmt.Println("inside handler", id)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not ok",
		})
		return
	}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := cr.Client.AddAddress(c.Request.Context(), body, id.(uint))
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
