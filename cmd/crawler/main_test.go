package crawler_test

import (
	"testing"
	"time"

	"github.com/conceptcodes/webcrawler-go/cmd/crawler"
)

func TestCrawlerQueueInit(t *testing.T) {
	// Arrange
	c := crawler.New("http://example.com", 5*time.Second)

	// Assert
	if !c.Queue.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it is not.")
	}
}

func TestCrawlerAddUrl(t *testing.T) {
	// Arrange
	c := crawler.New("http://example.com", 5*time.Second)
	url := "http://example.com/test"

	// Act
	c.AddUrl(url)

	// Assert
	if c.Queue.IsEmpty() {
		t.Errorf("Expected queue to not be empty, but it is.")
	}
	if c.Queue.Len() != 1 {
		t.Errorf("Expected queue size to be 1, but got %d.", c.Queue.Len())
	}
	if c.Queue.Dequeue() != url {
		t.Errorf("Expected dequeued URL to be %s, but got a different URL.", url)
	}
}

func TestCrawlerProcessQueue(t *testing.T) {
	// Arrange
	c := crawler.New("http://example.com", 5*time.Second)
	c.AddUrl("http://example.com/test1")
	c.AddUrl("http://example.com/test2")

	// Act
	c.ProcessQueue()

	// Assert
	if !c.Queue.IsEmpty() {
		t.Errorf("Expected queue to be empty after processing, but it is not.")
	}
}
