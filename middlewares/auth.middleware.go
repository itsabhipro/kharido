package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsabhipro/kharido/utils"
)

func AuthMiddleware(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")
	authuser, err := utils.VerifyAuthToken(authToken)
	if err != nil {
		c.Set("user", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invailid token"})
	}
	c.Set("user", authuser)
	c.Next()
}
