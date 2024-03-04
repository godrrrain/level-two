package handler

import (
	"dev11/storage"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type message struct {
	Message string `json:"message"`
}

type createEvent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`
}

// CreateEventHandler обрабатывает запросы на создание события
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "error: Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var input createEvent

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "error: Invalid request body", http.StatusBadRequest)
		return
	}

	start, err := time.Parse(time.RFC3339, input.Start)
	if err != nil {
		http.Error(w, "error: Invalid start time format", http.StatusBadRequest)
		return
	}
	end, err := time.Parse(time.RFC3339, input.End)
	if err != nil {
		http.Error(w, "error: Invalid end time format", http.StatusBadRequest)
		return
	}

	event, err := storage.Cache.CreateEvent(input.Title, input.Description, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

// GetEventsForDayHandler обрабатывает запросы на получение событий за день.
func GetEventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	day, err := parseDateQuery(r, "day")
	if err != nil {
		http.Error(w, "error: Invalid day query parameter", http.StatusBadRequest)
		return
	}

	events := storage.Cache.GetEventsForDay(day)
	if len(events) == 0 {
		json.NewEncoder(w).Encode(message{
			Message: "пусто",
		})
	} else {
		json.NewEncoder(w).Encode(events)
	}
}

// GetEventsForWeekHandler обрабатывает запросы на получение событий за неделю.
func GetEventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	startOfWeek, err := parseDateQuery(r, "week")
	if err != nil {
		http.Error(w, "error: Invalid week query parameter", http.StatusBadRequest)
		return
	}

	events := storage.Cache.GetEventsForWeek(startOfWeek)
	if len(events) == 0 {
		json.NewEncoder(w).Encode(message{
			Message: "пусто",
		})
	} else {
		json.NewEncoder(w).Encode(events)
	}
}

// GetEventsForMonthHandler обрабатывает запросы на получение событий за месяц.
func GetEventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	year, month, err := parseMonthQuery(r, "month")
	if err != nil {
		http.Error(w, "error: Invalid month query parameter", http.StatusBadRequest)
		return
	}

	events := storage.Cache.GetEventsForMonth(year, month)
	if len(events) == 0 {
		json.NewEncoder(w).Encode(message{
			Message: "пусто",
		})
	} else {
		json.NewEncoder(w).Encode(events)
	}
}

// UpdateEventHandler обрабатывает запросы на обновление события.
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "error: Invalid method", http.StatusMethodNotAllowed)
		return
	}

	updatedEvent, err := parseAndUpdateEvent(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = storage.Cache.UpdateEvent(updatedEvent.ID, updatedEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": "Event updated"})
}

// DeleteEventHandler обрабатывает запросы на удаление события.
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "error: Invalid method", http.StatusMethodNotAllowed)
		return
	}

	eventIDStr := r.URL.Query().Get("id")
	if eventIDStr == "" {
		http.Error(w, "error: Event ID is required", http.StatusBadRequest)
		return
	}

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		http.Error(w, "error: Invalid event ID", http.StatusBadRequest)
		return
	}

	err = storage.Cache.DeleteEvent(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": "Event deleted"})
}

func parseAndUpdateEvent(r *http.Request) (storage.Event, error) {
	var updatedEvent storage.Event
	err := json.NewDecoder(r.Body).Decode(&updatedEvent)
	if err != nil {
		return storage.Event{}, err
	}
	return updatedEvent, nil
}

func parseDateQuery(r *http.Request, queryParam string) (time.Time, error) {
	dateStr := r.URL.Query().Get(queryParam)
	if dateStr == "" {
		return time.Time{}, errors.New("error: query parameter missing")
	}
	return time.Parse("2006-01-02", dateStr)
}

func parseMonthQuery(r *http.Request, queryParam string) (int, time.Month, error) {
	monthStr := r.URL.Query().Get(queryParam)
	if monthStr == "" {
		return 0, 0, errors.New("error: query parameter missing")
	}
	date, err := time.Parse("2006-01", monthStr)
	if err != nil {
		return 0, 0, err
	}
	return date.Year(), date.Month(), nil
}
