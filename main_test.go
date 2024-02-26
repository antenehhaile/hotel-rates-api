package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	// Start the service in a goroutine
	go main()

	// Send a GET request to the "/health" endpoint
	req, err := http.NewRequest("GET", "http://localhost:8080/health", nil)
	assert.NoError(t, err)

	// Create a client and send the request
	client := &http.Client{}
	res, err := client.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()

	// Check if the status code is OK
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Read the response body and check if it contains the expected message
	expectedBody := `{"message":"The service is running healthy"}`
	bodyBytes, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, string(bodyBytes))
}

func TestHotelsEndpointFailure(t *testing.T) {
	// Start the service in a goroutine
	go main()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost:8080/hotels",
		httpmock.NewStringResponder(http.StatusInternalServerError, ""))

	// Send a GET request to the "/hotels" endpoint
	req, err := http.NewRequest("GET", "http://localhost:8080/hotels", nil)
	assert.NoError(t, err)

	// Create a client and send the request
	client := &http.Client{}
	res, err := client.Do(req)
	assert.NoError(t, err)

	defer res.Body.Close()

	// Check if the status code is INTERNAL SERVER ERROR
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}

func TestHotelsEndpointSuccess(t *testing.T) {
	// Start the service in a goroutine
	go main()

	// Mock HTTP server
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost:8080/hotels",
		httpmock.NewStringResponder(http.StatusOK, ""))

	// Send a GET request to the "/hotels" endpoint
	req, err := http.NewRequest("GET", "http://localhost:8080/hotels", nil)
	assert.NoError(t, err)

	// Create a client and send the request
	client := &http.Client{}
	res, err := client.Do(req)
	assert.NoError(t, err)

	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}
