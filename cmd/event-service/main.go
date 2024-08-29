package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chigaji/event-booking-system/cmd/event-service/db"
	"github.com/chigaji/event-booking-system/cmd/event-service/handlers"
	"github.com/gorilla/mux"
)

func main() {

	//initialize db
	db.InitDB()

	//router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/events", handlers.CreateEventHandler).Methods("POST")
	r.HandleFunc("/events", handlers.GetEventHander).Methods("GET")

	fmt.Println("Running server on port 8081")

	log.Fatal(http.ListenAndServe(":8081", r))
}
