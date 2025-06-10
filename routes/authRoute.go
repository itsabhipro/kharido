package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itsabhipro/kharido/controllers"
	"github.com/itsabhipro/kharido/middlewares"
)

func AuthRoute(ser *gin.RouterGroup) {
	ser.POST("/login", controllers.Login)
	ser.POST("/register", controllers.Register)
	ser.POST("/logout", controllers.Logout)
	ser.GET("/profile/:id", middlewares.AuthMiddleware, controllers.GetProfile)
	ser.PUT("/profile/:id", middlewares.AuthMiddleware, controllers.UpdateProfile)
}
