// routes.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	// Handle the index route
	router.GET("/", showIndexPage)
}

func showIndexPage(c *gin.Context) {
	// call the html method of the context to render a template
	c.HTML(
		// Set the HTTP status to 200(OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that page uses(in this case 'title')
		gin.H{
			"title": "HOME Page",
		},
	)
}
