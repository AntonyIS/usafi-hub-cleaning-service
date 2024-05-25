package domain

import "time"

type Service struct {
	ServiceId    string    `json:"service_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	PricePerHour int       `json:"price_per_hour"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Request struct {
	RequestId     string    `json:"request_id"`
	ClientId      string    `json:"client_id"`
	CleanerId     string    `json:"cleaner_id"`
	ServiceId     string    `json:"service_id"`
	RequestedDate time.Time `json:"requested_date"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Reviews struct {
	ReviewId  string    `json:"review_id"`
	RequestId string    `json:"request_id"`
	ClientId  string    `json:"client_id"`
	CleanerId string    `json:"cleaner_id"`
	Rating    string    `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
