package main

import (
	"fmt"
	"strings"
	"time"
)

type Event struct {
	EventID int       `json:"event_id"`
	UserID  int       `json:"user_id"`
	Title   string    `json:"title"`
	Info    string    `json:"info"`
	Date    time.Time `json:"date"`
}

type EventsMap struct {
	Events map[string]Event
}

func NewEvents() *EventsMap {
	return &EventsMap{
		Events: make(map[string]Event),
	}
}

func (e *EventsMap) CreateEvent(event *Event) error {
	UidEid := fmt.Sprintf("%d:%d", event.UserID, event.EventID)
	if _, ok := e.Events[UidEid]; ok {
		return fmt.Errorf("error: event with this ID already exists")
	}
	e.Events[UidEid] = *event
	return nil
}

func (e *EventsMap) UpdateEvent(userID, eventID int, newEvent *Event) error {
	UidEid := fmt.Sprintf("%d:%d", userID, eventID)
	if _, ok := e.Events[UidEid]; !ok {
		return fmt.Errorf("error: event with this ID does not exist")
	}
	e.Events[UidEid] = *newEvent
	return nil
}

func (e *EventsMap) DeleteEvent(userID, eventID int) error {
	UidEid := fmt.Sprintf("%d:%d", userID, eventID)
	if _, ok := e.Events[UidEid]; !ok {
		return fmt.Errorf("error: event with this ID does not exist")
	}
	delete(e.Events, UidEid)
	return nil
}

func (e *EventsMap) EventsForDay(userId string, day time.Time) ([]Event, error) {
	var events []Event
	for k, v := range e.Events {
		keyId := strings.Split(k, ":")[0]
		if userId == keyId {
			if v.Date.Day() == day.Day() {
				events = append(events, v)
			}
		}
	}

	if len(events) > 0 {
		return events, nil
	} else {
		return nil, fmt.Errorf("error: no events for this date ")
	}
}
func (e *EventsMap) EventsForWeek(userId string, week int) ([]Event, error) {
	var events []Event
	for k, v := range e.Events {
		keyId := strings.Split(k, ":")[0]
		if userId == keyId {
			_, w := v.Date.ISOWeek()
			if w == week {
				events = append(events, v)
			}
		}
	}

	if len(events) > 0 {
		return events, nil
	} else {
		return nil, fmt.Errorf("error: no events for this date ")
	}
}
func (e *EventsMap) EventsForMonth(userId string, month time.Month) ([]Event, error) {
	var events []Event
	for k, v := range e.Events {
		keyId := strings.Split(k, ":")[0]
		if userId == keyId {
			if v.Date.Month() == month {
				events = append(events, v)
			}
		}
	}

	if len(events) > 0 {
		return events, nil
	} else {
		return nil, fmt.Errorf("error: no events for this date ")
	}
}
