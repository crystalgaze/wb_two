package main

import (
	"fmt"
	"log"
	"net/http"
)

func runServer(port int, h *Handler) error {

	http.HandleFunc("/create_event", h.CreateEventHandler)
	http.HandleFunc("/update_event", h.UpdateEventHandler)
	http.HandleFunc("/delete_event", h.DeleteEventHandler)
	http.HandleFunc("/events_for_day/", h.EventsForDayHandler)
	http.HandleFunc("/events_for_week/", h.EventsForWeekHandler)
	http.HandleFunc("/events_for_month/", h.EventsForMonthHandler)

	log.Println("Starting http server:", fmt.Sprintf("http://127.0.0.1:%v/", port))
	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
