package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsabhipro/kharido/models"
	"github.com/itsabhipro/kharido/utils"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "this is login end point"})
}

func Register(c *gin.Context) {
	var userPayload models.User
	err := c.ShouldBindBodyWithJSON(&userPayload)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invailid data."})
		return
	}
	authToken, err := utils.GenerateAuthToken(&userPayload)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "token": authToken})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "this is logout endpoint"})
}

func GetProfile(c *gin.Context) {
	id, _ := c.Params.Get("id")
	c.JSON(http.StatusOK, gin.H{"message": "this is profile end", "id": id})
}

func UpdateProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "this is update profile end"})
}
