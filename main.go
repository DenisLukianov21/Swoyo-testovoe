package main

import (
	c "shortener-api/iternal/shortener"

	"github.com/gin-gonic/gin"
)

// main is the entry point of the application.
//
// It initializes a Gin router with default middleware and registers two routes:
// - POST / for the CreateShortUrl handler
// - GET /:url for the GetRawUrl handler
//
// The application is then started and listens on localhost port 8080.
func main() {
	router := gin.Default()
	router.POST("", c.CreateShortUrl)
	router.GET("/:url", c.GetRawUrl)
	router.Run("localhost:8080")
}
