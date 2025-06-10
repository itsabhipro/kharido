package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsabhipro/kharido/routes"
)

type application struct {
	port string
}

func (app *application) serve() {
	server := gin.Default()
	authRout := server.Group("/api/v1")
	routes.AuthRoute(authRout)
	server.Run(app.port)
}
