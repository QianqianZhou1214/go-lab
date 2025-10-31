package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Default() returns an Engine instance with Logger and Recovery (when panic)

	// Define route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Run() attaches the router to a http.Server and starts listening and serving HTTP requests.
	r.Run()
}
