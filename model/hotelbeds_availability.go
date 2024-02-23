package model

type AuditData struct {
	ProcessTime string `json:"processTime"`
	Timestamp   string `json:"timestamp"`
	RequestHost string `json:"requestHost"`
	ServerId    string `json:"serverId"`
	Environment string `json:"environment"`
	Release     string `json:"release"`
	Token       string `json:"token"`
	Internal    string `json:"internal"`
}

type Rate struct {
	RateKey              string `json:"rateKey"`
	RateClass            string `json:"rateClass"`
	RateType             string `json:"rateType"`
	Net                  string `json:"net"`
	Allotment            int    `json:"allotment"`
	RateCommentsId       string `json:"rateCommentsId"`
	PaymentType          string `json:"paymentType"`
	Packaging            bool   `json:"packaging"`
	BoardCode            string `json:"boardCode"`
	BoardName            string `json:"boardName"`
	CancellationPolicies []struct {
		Amount string `json:"amount"`
		From   string `json:"from"`
	} `json:"cancellationPolicies"`
	Taxes struct {
		Taxes []struct {
			Included       bool   `json:"included"`
			Amount         string `json:"amount"`
			Currency       string `json:"currency"`
			ClientAmount   string `json:"clientAmount"`
			ClientCurrency string `json:"clientCurrency"`
		} `json:"taxes"`
		AllIncluded bool `json:"allIncluded"`
	} `json:"taxes"`
	Rooms      int `json:"rooms"`
	Adults     int `json:"adults"`
	Children   int `json:"children"`
	Promotions []struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Remark string `json:"remark"`
	} `json:"promotions"`
}

type HotelbedsRoom struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Rates []Rate `json:"rates"`
}

type Hotel struct {
	Code            int             `json:"code"`
	Name            string          `json:"name"`
	ExclusiveDeal   int             `json:"exclusiveDeal"`
	CategoryCode    string          `json:"categoryCode"`
	CategoryName    string          `json:"categoryName"`
	DestinationCode string          `json:"destinationCode"`
	DestinationName string          `json:"destinationName"`
	ZoneCode        int             `json:"zoneCode"`
	ZoneName        string          `json:"zoneName"`
	Latitude        string          `json:"latitude"`
	Longitude       string          `json:"longitude"`
	Rooms           []HotelbedsRoom `json:"rooms"`
	MinRate         string          `json:"minRate,omitempty"`
	MaxRate         string          `json:"maxRate"`
	Currency        string          `json:"currency"`
}

type Hotels struct {
	Hotels   []Hotel `json:"hotels"`
	CheckIn  string  `json:"checkIn"`
	Total    int     `json:"total"`
	CheckOut string  `json:"checkOut"`
}

type HotelAvailability struct {
	AuditData AuditData `json:"auditData"`
	Hotels    Hotels    `json:"hotels"`
}
