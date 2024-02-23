package model

import (
	"time"
)

type RoomGuest struct {
	AdultCount int `json:"adultCount"`
	ChildCount int `json:"childCount"`
}

type Board struct {
	Boards []interface{} `json:"boards"`
}

type IncludedBoard struct {
	BoardId          string `json:"boardId"`
	BoardDescription string `json:"boardDescription"`
}

type RoomRate struct {
	InitialPrice         string   `json:"initialPrice"`
	Price                string   `json:"price"`
	InitialPricePerNight []string `json:"initialPricePerNight"`
	PricePerNight        []string `json:"pricePerNight"`
}

type Room struct {
	RoomCode        string        `json:"roomCode"`
	AdultCount      int           `json:"adultCount"`
	ChildCount      int           `json:"childCount"`
	Boards          Board         `json:"boards"`
	IncludedBoard   IncludedBoard `json:"includedBoard"`
	RoomDescription string        `json:"roomDescription"`
	RoomRemarks     string        `json:"roomRemarks"`
	RoomRate        RoomRate      `json:"roomRate"`
}

type TaxesAndFeesList struct {
	TaxesAndFees []TaxesAndFee `json:"taxesAndFees"`
}

type TaxesAndFee struct {
	Included    bool   `json:"included"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
}

type CancelPoliciesInfos struct {
	CancelPolicyInfos []struct {
		CancelTime string `json:"cancelTime"`
		Amount     string `json:"amount"`
		Type       string `json:"type"`
	} `json:"cancelPolicyInfos"`
	HotelRemarks         []interface{} `json:"hotelRemarks"`
	CancellationPolicies []interface{} `json:"cancellationPolicies"`
	RefundableTag        string        `json:"refundableTag"`
}

type RateDetails struct {
	RateDetails []RateDetail `json:"rateDetails"`
	MinimumRate string       `json:"minRate"`
	MaximumRate string       `json:"maxRate"`
}

type RateDetail struct {
	RateDetailCode      string              `json:"rateDetailCode"`
	RoomsList           RoomsList           `json:"rooms"`
	TotalPrice          string              `json:"totalPrice"`
	Tax                 int                 `json:"tax"`
	HotelFees           string              `json:"hotelFees"`
	Remarks             string              `json:"remarks"`
	TaxesAndFeesList    TaxesAndFeesList    `json:"taxesAndFees"`
	CancelPoliciesInfos CancelPoliciesInfos `json:"cancelPoliciesInfos"`
}

type RoomsList struct {
	Rooms []Room `json:"rooms"`
}

type HotelInfo struct {
	HotelCode       int    `json:"hotelCode"`
	HotelName       string `json:"hotelName"`
	HotelAddress    string `json:"hotelAddress"`
	HotelPictureUrl string `json:"hotelPictureUrl"`
	Longitude       string `json:"longitude"`
	Latitude        string `json:"latitude"`
	StarRating      string `json:"starRating"`
}

type HotelResult struct {
	HotelInfo       HotelInfo   `json:"hotelInfo"`
	MinPrice        string      `json:"minPrice"`
	RateDetailsList RateDetails `json:"rateDetails"`
}

type ResponseStatus struct {
	StatusCode    string    `json:"statusCode"`
	StatusMessage string    `json:"StatusMessage"`
	RequestAt     time.Time `json:"requestAt"`
	ResponseAt    time.Time `json:"responseAt"`
}

type City struct {
	CityCode    int    `json:"cityCode"`
	CityName    string `json:"cityName"`
	CountryName string `json:"countryName"`
}

type SupplierResponse struct {
	ResponseStatus   ResponseStatus `json:"responseStatus"`
	SessionId        string         `json:"sessionId"`
	City             City           `json:"city"`
	CheckInDate      string         `json:"checkInDate"`
	CheckOutDate     string         `json:"checkOutDate"`
	Currency         string         `json:"currency"`
	RoomGuestsList   RoomGuestsList `json:"roomGuests"`
	HotelResultsList HotelResults   `json:"hotelResults"`
}

type HotelResults struct {
	HotelResults []HotelResult `json:"hotelResults"`
}
type SupplierRequest struct {
	CheckInDate      string         `json:"checkInDate"`
	CheckOutDate     string         `json:"checkOutDate"`
	HotelCodes       []int          `json:"hotelCodes"`
	RoomGuestsList   RoomGuestsList `json:"roomGuests"`
	GuestNationality string         `json:"guestNationality"`
	Currency         string         `json:"currency"`
	LanguageCode     string         `json:"languageCode"`
	Timeout          string         `json:"timeout"`
}

type RoomGuestsList struct {
	RoomGuests []RoomGuest `json:"roomGuests"`
}

type ResponseData struct {
	Data     []Data   `json:"data"`
	Supplier Supplier `json:"supplier"`
}

type Data struct {
	HotelId  string `json:"hotelId"`
	Currency string `json:"currency"`
	Price    string `json:"price"`
}

type Supplier struct {
	Request  SupplierRequest  `json:"request"`
	Response SupplierResponse `json:"response"`
}
