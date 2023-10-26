package model

//FLIGHT TYPE

type FlightType int

const (
	Commercial FlightType = iota
	Military
	Cargo
)

type FlightTypeRequest struct {
	Type                FlightType `json:"type" gorm:"not null"`
	FlightModel         string     `json:"flight_model" gorm:"not null"`
	Description         string     `json:"description" gorm:"not null"`
	ManufacturerName    string     `json:"maunufacturer_name" gorm:"not null"`
	ManufacturerCountry string     `json:"manufacturer_country" gorm:"not null"`
	MaxDistance         string     `json:"max_distance" gorm:"not null"`
	CruiseSpeed         string     `json:"cruise_speed" gorm:"not null"`
}

// AIRLINE COMPANY

type AirlineCompanyRequest struct {
	AirlineName         string `json:"name"`
	CompanyAddress      string `json:"company_address"`
	PhoneNumber         string `json:"phone_number"`
	Email               string `json:"email"`
	AirlineCode         string `json:"airline_code"`
	AirlineLogoLink     string `json:"airline_logo_link"`
	SupportDocumentLink string `json:"support_documents_link"`
}

// AIRLINE SEATS

type AirlineSeatRequest struct {
	AirlineFleetID     uint `json:"airline_fleet_id"`
	EconomySeatNumber  uint `json:"economy_seat_no"`
	BuisinesSeatNumber uint `json:"buisines_seat_no"`
}

// AIRLINE BAGGAGE

// type AirlineBaggage

// AIRLINE CANCELLATION
