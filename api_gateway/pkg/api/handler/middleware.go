package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cr *AuthHandler) AuthorizationMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(role + "-token")
		if err != nil || tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Needs to login",
			})
			return
		}
		res, err1 := cr.Client.AuthorizationMiddleware(tokenString)
		if err1 != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization failed" + res.Error,
			})
			return
		}
		c.Set(role+"-id", uint(res.UserId))
		c.Next()
	}
}
