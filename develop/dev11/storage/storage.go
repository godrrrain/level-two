package storage

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Cache является глобальным
var Cache = NewEventsCache()

// Event представляет событие в календаре.
type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
}

// EventsCache представляет кэш событий.
type EventsCache struct {
	sync.RWMutex
	events map[int]Event
	nextID int
}

// NewEventsCache создает новый кэш событий.
func NewEventsCache() *EventsCache {
	return &EventsCache{
		events: make(map[int]Event),
		nextID: 1, // Начинаем ID с 1
	}
}

// CreateEvent добавляет новое событие и возвращает его.
func (c *EventsCache) CreateEvent(title, description string, start, end time.Time) (Event, error) {
	c.Lock()
	defer c.Unlock()

	if start.After(end) {
		return Event{}, errors.New("start time must be before end time")
	}

	event := Event{
		ID:          c.nextID,
		Title:       title,
		Description: description,
		Start:       start,
		End:         end,
	}

	c.events[event.ID] = event
	c.nextID++

	return event, nil
}

// UpdateEvent обновляет существующее событие
func (c *EventsCache) UpdateEvent(eventID int, updatedEvent Event) error {
	c.Lock()
	defer c.Unlock()

	if _, exists := c.events[eventID]; !exists {
		return fmt.Errorf("event with ID %d does not exist", eventID)
	}

	c.events[eventID] = updatedEvent
	return nil
}

// DeleteEvent удаляет событие с заданным ID.
func (c *EventsCache) DeleteEvent(eventID int) error {
	c.Lock()
	defer c.Unlock()

	if _, exists := c.events[eventID]; !exists {
		return fmt.Errorf("event with ID %d does not exist", eventID)
	}

	delete(c.events, eventID)
	return nil
}

// GetEventsForDay возвращает все события для указанного дня.
func (c *EventsCache) GetEventsForDay(day time.Time) []Event {
	c.RLock()
	defer c.RUnlock()

	var eventsForDay []Event
	for _, event := range c.events {
		if sameDay(event.Start, day) {
			eventsForDay = append(eventsForDay, event)
		}
	}
	return eventsForDay
}

// GetEventsForWeek возвращает все события для указанной недели.
func (c *EventsCache) GetEventsForWeek(startOfWeek time.Time) []Event {
	c.RLock()
	defer c.RUnlock()

	var eventsForWeek []Event
	endOfWeek := startOfWeek.AddDate(0, 0, 7)
	for _, event := range c.events {
		if event.Start.After(startOfWeek) && event.Start.Before(endOfWeek) {
			eventsForWeek = append(eventsForWeek, event)
		}
	}
	return eventsForWeek
}

// GetEventsForMonth возвращает все события для указанного месяца.
func (c *EventsCache) GetEventsForMonth(year int, month time.Month) []Event {
	c.RLock()
	defer c.RUnlock()

	var eventsForMonth []Event
	for _, event := range c.events {
		if event.Start.Year() == year && event.Start.Month() == month {
			eventsForMonth = append(eventsForMonth, event)
		}
	}
	return eventsForMonth
}

func sameDay(a, b time.Time) bool {
	y1, m1, d1 := a.Date()
	y2, m2, d2 := b.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
