package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hotel-rates-api/model"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHotelCheapRatesSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock API key and secret
	apiKey := "mock-api-key"
	secret := "mock-secret"

	// Mock Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Mock HTTP request
	req, err := http.NewRequest(http.MethodGet, "/hotels", nil)
	require.NoError(t, err)

	// Set up request queries
	queries := map[string]string{
		"checkin":          "2024-06-15",
		"checkout":         "2024-06-16",
		"currency":         "USD",
		"guestNationality": "US",
		"hotelIds":         "264",
		"occupancies":      `[{"rooms":1, "adults": 2}]`,
	}

	q := req.URL.Query()
	for key, value := range queries {
		q.Set(key, value)
	}
	req.URL.RawQuery = q.Encode()

	// Set the mock request to Gin context's request
	c.Request = req

	// Mock HotelAvailability
	hotelbedsApiResponseBytes, err := ioutil.ReadFile("test-data/hotelbeds_api_success_api_reposnse.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Step 2: Unmarshal JSON data into struct
	var mockResponseHotelAvailability model.HotelAvailability
	if err := json.Unmarshal(hotelbedsApiResponseBytes, &mockResponseHotelAvailability); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	responseBytesHotelAvailability, err := json.Marshal(mockResponseHotelAvailability)
	require.NoError(t, err)

	ratesApiResponseBytes, err := ioutil.ReadFile("test-data/rates_api_success_reposnse.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Mock HTTP server
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://api.test.hotelbeds.com/hotel-api/1.0/hotels",
		httpmock.NewStringResponder(http.StatusOK, string(responseBytesHotelAvailability)))

	// Initialize hotel service
	hotelService := NewHotelService()

	// Call the function being tested
	err = hotelService.GetHotelCheapRates(c, apiKey, secret)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert the response
	var expectedResponse model.ResponseData
	err = json.Unmarshal(w.Body.Bytes(), &expectedResponse)
	require.NoError(t, err)

	modifiedActual, err := removeField(string(w.Body.Bytes()), "supplier", "response", "responseStatus")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	modifiedExpected, err := removeField(string(ratesApiResponseBytes), "supplier", "response", "responseStatus")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data1 map[string]interface{}
	json.Unmarshal([]byte(modifiedActual), &data1)

	var data2 map[string]interface{}
	json.Unmarshal([]byte(modifiedExpected), &data2)

	assert.Equal(t, data1, data2)
}

func TestGetHotelCheapRatesHotelBedsApiFailure(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock API key and secret
	apiKey := "mock-api-key"
	secret := "mock-secret"

	// Mock Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Mock HTTP request
	req, err := http.NewRequest(http.MethodGet, "/hotels", nil)
	require.NoError(t, err)

	// Set up request queries
	queries := map[string]string{
		"checkin":          "2024-06-15",
		"checkout":         "2024-06-16",
		"currency":         "USD",
		"guestNationality": "US",
		"hotelIds":         "264",
		"occupancies":      `[{"rooms":1, "adults": 2}]`,
	}

	q := req.URL.Query()
	for key, value := range queries {
		q.Set(key, value)
	}
	req.URL.RawQuery = q.Encode()

	// Set the mock request to Gin context's request
	c.Request = req

	// Mock HTTP server
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://api.test.hotelbeds.com/hotel-api/1.0/hotels",
		httpmock.NewStringResponder(http.StatusBadRequest, "Bad Request"))

	// Initialize hotel service
	hotelService := NewHotelService()

	// Call the function being tested
	err = hotelService.GetHotelCheapRates(c, apiKey, secret)

	// Assert that no error occurred
	assert.Error(t, err)
}

func removeField(jsonString string, fieldPath ...string) (string, error) {
	// Unmarshal JSON string into a map
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return "", err
	}

	// Traverse nested structure to reach the specified field
	currentMap := data
	for i, path := range fieldPath {
		if i == len(fieldPath)-1 {
			// Delete the field at the last path segment
			delete(currentMap, path)
			break
		}
		// Check if the current path exists and is a nested map
		value, ok := currentMap[path].(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("path '%s' does not exist or is not a nested map", path)
		}
		currentMap = value
	}

	// Marshal the modified map back to JSON string
	modifiedJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(modifiedJSON), nil
}
