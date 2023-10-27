package DTO

//FLIGHT TYPE

type FlightType int

const (
	Commercial FlightType = iota
	Military
	Cargo
)

type FlightTypeRequest struct {
	Type                FlightType `json:"type" validate:"max=2,required"`
	FlightModel         string     `json:"flight_model" validate:"max=5,required"`
	Description         string     `json:"description" validate:"max=150,required"`
	ManufacturerName    string     `json:"maunufacturer_name" validate:"max=150,required"`
	ManufacturerCountry string     `json:"manufacturer_country" validate:"max=150,required"`
	MaxDistance         int32      `json:"max_distance" validate:"max=1500,required"`
	CruiseSpeed         int32      `json:"cruise_speed" validate:"max=1500,required"`
}

// AIRLINE COMPANY

type AirlineCompanyRequest struct {
	AirlineName         string `json:"name" validate:"required"`
	CompanyAddress      string `json:"company_address" validate:"max=100,required"`
	PhoneNumber         string `json:"phone_number" validate:"number,required"`
	Email               string `json:"email" validate:"email,required"`
	AirlineCode         string `json:"airline_code" validate:"max=5required"`
	AirlineLogoLink     string `json:"airline_logo_link" validate:"url,required"`
	SupportDocumentLink string `json:"support_documents_link" validate:"url,required"`
}

// AIRLINE SEATS

type AirlineSeatRequest struct {
	AirlineId           int `json:"airline_id" validate:"required,min=0"`
	EconomySeatNumber   int `json:"economy_seat_no" validate:"required,number"`
	BuisinesSeatNumber  int `json:"buisines_seat_no" validate:"required,number"`
	EconomySeatsPerRow  int `json:"economy_seats_per_row" validate:"required,number"`
	BuisinesSeatsPerRow int `json:"buisines_seats_per_row" validate:"required,number"`
}

// AIRLINE BAGGAGE

type Class int

const (
	Economy Class = iota
	Buisiness
)

type AirlineBaggageRequest struct {
	AirlineId           int    `json:"airline_id" validate:"required,number"`
	FareClass           int    `json:"class" validate:"required"`
	CabinAllowedWeight  int    `json:"cabin_allowed_weight" validate:"required,number"`
	CabinAllowedLength  int    `json:"cabin_allowed_length" validate:"required"`
	CabinAllowedBreadth int    `json:"cabin_allowed_breadth" validate:"required"`
	CabinAllowedHeight  int    `json:"cabin_allowed_height" validate:"required"`
	HandAllowedWeight   int    `json:"hand_allowed_weight" validate:"required,number"`
	HandAllowedLength   int    `json:"hand_allowed_length" validate:"required"`
	HandAllowedBreadth  int    `json:"hand_allowed_breadth" validate:"required"`
	HandAllowedHeight   int    `json:"hand_allowed_height" validate:"required"`
	FeeExtraPerKGCabin  int    `json:"fee_for_extra_kg_cabin" validate:"required,number"`
	FeeExtraPerKGHand   int    `json:"fee_for_extra_kg_hand" validate:"required,number"`
	Restrictions        string `json:"restrictions"`
}

// AIRLINE CANCELLATION

type AirlineCancellationRequest struct {
	AirlineId                  int  `json:"airline_id" validate:"required,number"`
	FareClass                  int  `json:"class" validate:"required"`
	CancellationDeadlineBefore int `json:"cancellation_deadline_before_hours" validate:"required"`
	CancellationPercentage     int  `json:"cancellation_percentage" validate:"required"`
	Refundable                 bool `json:"refundable" validate:"required,boolean"`
}
