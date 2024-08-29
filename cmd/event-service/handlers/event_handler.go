package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chigaji/event-booking-system/cmd/event-service/db"
	"github.com/chigaji/event-booking-system/cmd/event-service/models"
)

/*
*
Handles event creation
*/
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	// db.InitDB()
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)

	fmt.Println(event)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()
	event.AvailableTickets = event.TotalTickets

	fmt.Println("event after: ", event)

	query := `INSERT INTO events (title, description, total_tickets, available_tickets, event_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	fmt.Println("+++herer done 1")
	err = db.DB.QueryRow(query, event.Title, event.Description, event.TotalTickets, event.AvailableTickets, event.EventDate, event.CreatedAt, event.UpdatedAt).Scan(&event.ID)
	fmt.Println("+++herer done 2")
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}
	fmt.Println("+++herer done 3")
	w.WriteHeader(http.StatusCreated)
	fmt.Println("+++herer done 4")
	json.NewEncoder(w).Encode(event)

}

/*
*
Fetches event details
*/
func GetEventHander(w http.ResponseWriter, r *http.Request) {
	eventID := r.URL.Query().Get("id")

	if eventID == "" {
		http.Error(w, "missing event ID", http.StatusBadRequest)
		return
	}

	var event models.Event
	query := `SELECT id, time, description, total_tickets, available_tickets, event_date, created_at, updated_at
	FROM events WHERE id = $1`
	err := db.DB.QueryRow(query, eventID).Scan(&event)
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}
