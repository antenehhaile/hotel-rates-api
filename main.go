package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hotel-rates-api/service"
)

func main() {
	// apiKey := os.Getenv("API_KEY")
	// secret := os.Getenv("SECRET")
	fmt.Println("API is running ...")
	apiKey := "454754dfa49008c9539f0d082cd488b2"
	secret := "aebd06614b"
	r := gin.Default()

	// Initialize hotel service
	hotelService := service.NewHotelService()

	// Define routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "The service is running healthy"})
	})

	// Define routes
	r.GET("/hotels", func(c *gin.Context) {
		err := hotelService.GetHotelCheapRates(c, apiKey, secret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	})

	// Run the server on port 8080
	r.Run(":8080")
}
