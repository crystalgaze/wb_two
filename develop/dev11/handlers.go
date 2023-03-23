package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ResultResponse struct {
	Result []Event `json:"result"`
	Err    string  `json:"error"`
}

type Handler struct {
	events *EventsMap
}

func NewHandler() *Handler {
	return &Handler{
		events: NewEvents(),
	}
}

func (h *Handler) CreateEventHandler(w http.ResponseWriter, req *http.Request) {
	var event Event
	err := json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		errorResponse(w, err, http.StatusBadRequest)
		return
	}
	err = h.events.CreateEvent(&event)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	resultResponse(w, []Event{event})
	log.Printf("event for user %d with id %d created successfully", event.UserID, event.EventID)
}

func (h *Handler) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		errorResponse(w, err, http.StatusBadRequest)
		return
	}
	err = h.events.UpdateEvent(event.UserID, event.EventID, &event)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	resultResponse(w, []Event{event})
	log.Printf("event for user %d with id %d updated successfully", event.UserID, event.EventID)

}
func (h *Handler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.events.DeleteEvent(event.UserID, event.EventID)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	resultResponse(w, []Event{event})
	log.Printf("event for user %d with id %d deleted successfully", event.UserID, event.EventID)

}
func (h *Handler) EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02"
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	t, err := time.Parse(layout, date)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		errorResponse(w, err, http.StatusBadRequest)
		return
	}
	events, err := h.events.EventsForDay(userID, t)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	resultResponse(w, events)
	log.Printf("events for %d.%d.%d were shown successfully", t.Day(), t.Month(), t.Year())

}
func (h *Handler) EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02"
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	t, err := time.Parse(layout, date)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		errorResponse(w, err, http.StatusBadRequest)
		return
	}
	_, week := t.ISOWeek()

	events, err := h.events.EventsForWeek(userID, week)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	resultResponse(w, events)
	log.Printf("events for week %d year %d were shown successfully", week, t.Year())

}
func (h *Handler) EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02"
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	t, err := time.Parse(layout, date)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	events, err := h.events.EventsForMonth(userID, t.Month())
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		errorResponse(w, err, http.StatusServiceUnavailable)
	}
	resultResponse(w, events)
	log.Printf("events for month %d year %d were shown successfully", t.Month(), t.Year())
}

func resultResponse(w http.ResponseWriter, events []Event) {
	w.Header().Set("Content-Type", "application/json")
	result, _ := json.MarshalIndent(&ResultResponse{Result: events, Err: "-1"}, " ", "")
	_, err := w.Write(result)
	if err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
		return
	}
}

func errorResponse(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	jsonErr, _ := json.MarshalIndent(&ResultResponse{Result: nil, Err: err.Error()}, " ", " ")
	http.Error(w, string(jsonErr), status)
}
