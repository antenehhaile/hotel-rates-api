package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hotel-rates-api/service"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	secret := os.Getenv("SECRET")
	r := gin.Default()

	// Initialize hotel service
	hotelService := service.NewHotelService()

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
