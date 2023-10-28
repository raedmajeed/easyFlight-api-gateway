package DTO

//* FLIGHT TYPE

type FlightType int

const (
	Commercial FlightType = iota
	Military
	Cargo
)

type FlightTypeRequest struct {
	Type                FlightType `json:"type" validate:"max=2,required"`
	FlightModel         string     `json:"flight_model" validate:"max=5,required,alphanumunicode,upper"`
	Description         string     `json:"description" validate:"max=150,required"`
	ManufacturerName    string     `json:"maunufacturer_name" validate:"max=150,required,alpha"`
	ManufacturerCountry string     `json:"manufacturer_country" validate:"max=150,required,alpha"`
	MaxDistance         int32      `json:"max_distance" validate:"max=1500,required,numeric"`
	CruiseSpeed         int32      `json:"cruise_speed" validate:"max=1500,required,numeric"`
}

//* AIRLINE COMPANY

type AirlineCompanyRequest struct {
	AirlineName         string `json:"name" validate:"required"`
	CompanyAddress      string `json:"company_address" validate:"max=100,required"`
	PhoneNumber         string `json:"phone_number" validate:"phone,required"`
	Email               string `json:"email" validate:"email,required"`
	AirlineCode         string `json:"airline_code" validate:"max=5,required,alphanumeric,upper"`
	AirlineLogoLink     string `json:"airline_logo_link" validate:"url,required"`
	SupportDocumentLink string `json:"support_documents_link" validate:"url,required"`
}

//* AIRLINE SEATS

type AirlineSeatRequest struct {
	AirlineId           int `json:"airline_id" validate:"required,min=0"`
	EconomySeatNumber   int `json:"economy_seat_no" validate:"required,number"`
	BuisinesSeatNumber  int `json:"buisines_seat_no" validate:"required,number"`
	EconomySeatsPerRow  int `json:"economy_seats_per_row" validate:"required,number"`
	BuisinesSeatsPerRow int `json:"buisines_seats_per_row" validate:"required,number"`
}

//* AIRLINE BAGGAGE

type Class int

const (
	Economy Class = iota
	Buisiness
)

type AirlineBaggageRequest struct {
	AirlineId           int    `json:"airline_id" validate:"required,number"`
	FareClass           int    `json:"class" validate:"required"`
	CabinAllowedWeight  int    `json:"cabin_allowed_weight" validate:"required,numeric"`
	CabinAllowedLength  int    `json:"cabin_allowed_length" validate:"required,numeric"`
	CabinAllowedBreadth int    `json:"cabin_allowed_breadth" validate:"required,numeric"`
	CabinAllowedHeight  int    `json:"cabin_allowed_height" validate:"required,numeric"`
	HandAllowedWeight   int    `json:"hand_allowed_weight" validate:"required,number,numeric"`
	HandAllowedLength   int    `json:"hand_allowed_length" validate:"required,numeric"`
	HandAllowedBreadth  int    `json:"hand_allowed_breadth" validate:"required,numeric"`
	HandAllowedHeight   int    `json:"hand_allowed_height" validate:"required,numeric"`
	FeeExtraPerKGCabin  int    `json:"fee_for_extra_kg_cabin" validate:"required,numeric"`
	FeeExtraPerKGHand   int    `json:"fee_for_extra_kg_hand" validate:"required,numeric"`
	Restrictions        string `json:"restrictions"`
}

//* AIRLINE CANCELLATION

type AirlineCancellationRequest struct {
	AirlineId                  int  `json:"airline_id" validate:"required,number"`
	FareClass                  int  `json:"class" validate:"required"`
	CancellationDeadlineBefore int  `json:"cancellation_deadline_before_hours" validate:"required,numeric"`
	CancellationPercentage     int  `json:"cancellation_percentage" validate:"required"`
	Refundable                 bool `json:"refundable" validate:"required,boolean"`
}

//* FLIGHT FLEET

type FlightFleetRequest struct {
	SeatId             int `json:"seat_id" validate:"required,foreign_key:seats"`
	FlightType         int `json:"flight_type" validate:"required,foreign_key:flight_types"`
	BaggagePolicy      int `json:"baggage_policy" validate:"required,foreign_key:baggage_policies"`
	CancellationPolicy int `json:"cancellation_policy" validate:"required,foreign_key:cancellation_policies"`
	FlightCount        int `json:"flight_count" validate:"numeric"`
}
