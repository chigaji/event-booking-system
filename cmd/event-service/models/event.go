package models

import "time"

//	type Event struct {
//		ID               int       `json:"id"`
//		Title            string    `json:"title"`
//		Description      string    `json:"description"`
//		TotalTickets     int       `json:"total_tickets"`
//		AvailableTickets int       `json:"available_tickets"`
//		EventDate        time.Time `json:"event_date"`
//		CreatedAt        time.Time `json:"created_at"`
//		UpdatedAt        time.Time `json:"updated_at"`
//	}
type Event struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	TotalTickets     int       `json:"total_tickets"`
	AvailableTickets int       `json:"available_tickets"`
	EventDate        time.Time `json:"event_date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
