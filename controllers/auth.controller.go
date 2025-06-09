package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "this is login end point"})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "this is register endpoint"})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "this is logout endpoint"})
}
