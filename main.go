package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// start serving the application
	router.Run()
}
