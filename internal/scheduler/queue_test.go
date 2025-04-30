package scheduler_test

import (
	"testing"

	"github.com/conceptcodes/webcrawler-go/internal/scheduler"
)

func TestEnqueue(t *testing.T) {
	// Arrange
	q := scheduler.New()

	// Act
	item := "test"
	q.Enqueue(item)

	// Assert
	if q.Len() != 1 {
		t.Errorf("Expected queue length to be 1, got %d", q.Len())
	}
}

func TestDequeue(t *testing.T) {
	// Arrange
	q := scheduler.New()
	item := "test"
	q.Enqueue(item)

	// Act
	dequeuedItem := q.Dequeue()

	// Assert
	if dequeuedItem != item {
		t.Errorf("Expected dequeued item to be %s, got %s", item, dequeuedItem)
	}
	if q.Len() != 0 {
		t.Errorf("Expected queue length to be 0, got %d", q.Len())
	}
}
