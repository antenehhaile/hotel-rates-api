package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hotel-rates-api/model"
)

type HotelService interface {
	GetHotelCheapRates(c *gin.Context, apiKey, secret string) error
}

type hotelService struct {
}

func NewHotelService() HotelService {
	return &hotelService{}
}

func (s *hotelService) GetHotelCheapRates(c *gin.Context, apiKey, secret string) error {
	// Get the current time
	start := time.Now()
	ctx := c.Request.Context()
	// Retrieve the request deadline (timeout)
	deadline, _ := ctx.Deadline()

	// Retrieve request query params
	checkin := c.Query("checkin")
	checkout := c.Query("checkout")
	currency := c.Query("currency")
	guestNationality := c.Query("guestNationality")
	hotelIds := c.Query("hotelIds")
	occ := c.Query("occupancies")
	var occupancies []model.Occupancy
	if occ != "" {
		err := json.Unmarshal([]byte(c.Query("occupancies")), &occupancies)
		if err != nil {
			fmt.Println("Error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": occ,
			})
			return errors.New("internal server error")
		}
	}

	hotelCodes, err := StringToIntArray(hotelIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return err
	}

	requestBody := &model.RequestBody{
		Stay: model.Stay{
			CheckIn:  checkin,
			CheckOut: checkout,
		},
		Occupancies: occupancies,
		Hotels: model.RequestHotels{
			Hotel: hotelCodes,
		},
	}

	url := "https://api.test.hotelbeds.com/hotel-api/1.0/hotels"

	// Define the request body
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return err
	}

	// Create a new HTTP request with POST method and request body
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-key", apiKey)
	req.Header.Set("X-Signature", generateSignature(apiKey, secret))

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	//Check if the API response is 200. If not, return an error
	if resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return errors.New("internal server error")
	}
	defer resp.Body.Close()

	end := time.Now()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	var hotelAvailability model.HotelAvailability
	err = json.Unmarshal(body, &hotelAvailability)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	var cheapestRates model.ResponseData
	var roomGuests []model.RoomGuest
	for _, r := range requestBody.Occupancies {
		roomGuests = append(roomGuests, model.RoomGuest{AdultCount: r.Adults, ChildCount: r.Children})
	}

	var hotelResults []model.HotelResult

	var data []model.Data
	for _, h := range hotelAvailability.Hotels.Hotels {

		data = append(data, model.Data{
			HotelId:  strconv.Itoa(h.Code),
			Currency: h.Currency,
			Price:    h.MinRate,
		})

		var rateDetails []model.RateDetail

		for _, room := range h.Rooms {

			var rooms []model.Room
			var taxesAndFees []model.TaxesAndFee

			for i, rate := range room.Rates {
				rooms = append(rooms, model.Room{
					RoomCode:   room.Code + "|" + strconv.Itoa(i),
					AdultCount: rate.Adults,
					ChildCount: rate.Children,
					Boards:     model.Board{},
					IncludedBoard: model.IncludedBoard{
						BoardId:          rate.BoardCode,
						BoardDescription: rate.BoardName,
					},
					RoomDescription: "",
					RoomRemarks:     "",
					RoomRate: model.RoomRate{
						InitialPrice:         rate.Net,
						Price:                rate.Net,
						InitialPricePerNight: []string{rate.Net},
						PricePerNight:        []string{rate.Net},
					},
				})
				for _, t := range rate.Taxes.Taxes {
					taxesAndFees = append(taxesAndFees, model.TaxesAndFee{
						Included:    t.Included,
						Description: "",
						Amount:      t.Amount,
						Currency:    t.Currency,
					})
				}
			}
			rateDetails = append(rateDetails, model.RateDetail{
				RateDetailCode: room.Code,
				RoomsList: model.RoomsList{
					Rooms: rooms,
				},
				TotalPrice: h.MinRate,
				TaxesAndFeesList: model.TaxesAndFeesList{
					TaxesAndFees: taxesAndFees,
				},
			})
		}

		//TODO: Some of the hotel infomation can be enriched by leveraging the hotels details endpoint ("/hotels/{hotelCodes}/details")
		hotelResults = append(hotelResults, model.HotelResult{
			HotelInfo: model.HotelInfo{
				HotelCode:       h.Code,
				HotelName:       h.Name,
				HotelAddress:    "",
				HotelPictureUrl: "",
				Longitude:       h.Longitude,
				Latitude:        h.Latitude,
				StarRating:      "",
			},
			MinPrice: h.MinRate,
			RateDetailsList: model.RateDetails{
				RateDetails: rateDetails,
				MinimumRate: h.MinRate,
				MaximumRate: h.MaxRate,
			},
		})

		cheapestRates = model.ResponseData{
			Data: data,
			Supplier: model.Supplier{
				Request: model.SupplierRequest{
					CheckInDate:  checkin,
					CheckOutDate: checkout,
					HotelCodes:   hotelCodes,
					RoomGuestsList: model.RoomGuestsList{
						RoomGuests: roomGuests,
					},
					GuestNationality: guestNationality,
					Currency:         currency,
					Timeout:          deadline.Format(time.RFC3339),
				},
				Response: model.SupplierResponse{
					ResponseStatus: model.ResponseStatus{
						StatusCode:    resp.Status,
						StatusMessage: "Success",
						RequestAt:     start,
						ResponseAt:    end,
					},
					SessionId:    hotelAvailability.AuditData.Internal,
					CheckInDate:  checkin,
					CheckOutDate: checkout,
					Currency:     currency,
					RoomGuestsList: model.RoomGuestsList{
						RoomGuests: roomGuests,
					},
					HotelResultsList: model.HotelResults{
						HotelResults: hotelResults,
					},
				},
			},
		}
	}

	// Return the response body
	c.JSON(http.StatusOK, cheapestRates)
	return nil
}

func StringToIntArray(str string) ([]int, error) {
	strs := strings.Split(str, ",")
	intArray := make([]int, len(strs))

	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		intArray[i] = num
	}

	return intArray, nil
}

func generateSignature(apiKey, secret string) string {
	// Get current timestamp in seconds
	timestamp := time.Now().Unix()

	// Concatenate apiKey, secret, and timestamp
	data := apiKey + secret + fmt.Sprintf("%d", timestamp)

	// Calculate SHA256 hash
	hash := sha256.New()
	hash.Write([]byte(data))
	signature := hex.EncodeToString(hash.Sum(nil))

	return signature
}
