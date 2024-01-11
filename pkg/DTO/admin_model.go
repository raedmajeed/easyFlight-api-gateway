package DTO

// *AIRPORT

type AirportRequest struct {
	AirportCode  string  `json:"airport_code" validate:"max=5,required,uppercase"`
	AirportName  string  `json:"airport_name" validate:"max=50,required,alphaspace"`
	City         string  `json:"city" validate:"max=20,required,alphaspace"`
	Country      string  `json:"country" validate:"max=20,required,alphaspace"`
	Region       string  `json:"region" validate:"max=50,required,alphaspace"`
	Latitude     float64 `json:"latitude" valdiate:"required,numeric,latitude"`
	Longitude    float64 `json:"longitude" validate:"required,numeric,longitude"`
	IATAFCSCode  string  `json:"iata_fcs_code" validate:"required,alphaspace,len=3"`
	ICAOCode     string  `json:"icao_code" validate:"required,alphaspace,len=4"`
	Website      string  `json:"website" validate:"required,url"`
	ContactEmail string  `json:"contact_email" validate:"required,emailcst"`
	ContactPhone string  `json:"contact_phone" validate:"required,phone"`
}

//* SCHEDULE

type ScheduleRequest struct {
	DepartureTime        string `json:"departure_time" validate:"required,time"`
	ArrivalTime          string `json:"arrival_time" validate:"required,time"`
	DepartureDate        string `json:"departure_date" validate:"required,date"`
	ArrivalDate          string `json:"arrival_date" validate:"required,date"`
	DepartureAirportCode string `json:"departure_airport_code" validate:"max=5,required,uppercase"`
	ArrivalAirportCode   string `json:"arrival_airport_code" validate:"max=5,required,uppercase"`
}

//* FLIGHT CHART REQUEST

type FlightChart struct {
	FlightFleetID int `json:"flight_fleet_id" validate:"required,numeric,min=0"`
	ScheduleID    int `json:"schedule_id" validate:"required,numeric,min=0"`
}

type FetchAirport struct {
	AirportCode string `json:"airport_code" validate:"required,max=5"`
}

type UpdateAirportRequest struct {
	Website      string `json:"website"`
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
}
