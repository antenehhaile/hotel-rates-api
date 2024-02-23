package model

type Occupancy struct {
	Rooms    int `json:"rooms"`
	Adults   int `json:"adults"`
	Children int `json:"children"`
}

type RequestBody struct {
	Stay        Stay          `json:"stay"`
	Occupancies []Occupancy   `json:"occupancies"`
	Hotels      RequestHotels `json:"hotels"`
	Language    string        `json:"language,omitempty"`
}

type Stay struct {
	CheckIn  string `json:"checkIn"`
	CheckOut string `json:"checkOut"`
}

type RequestHotels struct {
	Hotel []int `json:"hotel"`
}
