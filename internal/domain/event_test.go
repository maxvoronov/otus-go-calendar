package domain_test

import (
	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"testing"
	"time"
)

func TestNewEvent(t *testing.T) {
	title := "Demo Event"
	dateFrom := time.Now()
	dateTo := time.Now().Add(time.Hour)

	event := domain.NewEvent(title, dateFrom, dateTo)

	if event.ID.String() == "" {
		t.Fatal("Empty event ID")
	}

	if event.Title != title {
		t.Fatalf("Invalid event title: expect %s, got %s", title, event.Title)
	}

	if !event.DateFrom.Equal(dateFrom) {
		t.Fatalf("Invalid event date from: expect %s, got %s", dateFrom, event.DateFrom)
	}

	if !event.DateTo.Equal(dateTo) {
		t.Fatalf("Invalid event date to: expect %s, got %s", dateTo, event.DateTo)
	}
}
