package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itsabhipro/kharido/controllers"
)

func AuthRoute(ser *gin.RouterGroup) {
	ser.POST("/login", controllers.Login)
	ser.POST("/register", controllers.Register)
	ser.POST("/logout", controllers.Logout)
}
