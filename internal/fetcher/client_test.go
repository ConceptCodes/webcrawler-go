package fetcher_test

import (
	"testing"
	"time"

	"github.com/conceptcodes/webcrawler-go/internal/fetcher"
)

func TestFetchPageContents(t *testing.T) {
	// Arrange
	client := fetcher.New(5 * time.Second)

	// Act
	url := "http://example.com"
	content, err := client.FetchPageContents(url)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Assert
	if content == "" {
		t.Fatalf("expected non-empty content, got empty string")
	}
}

func TestFetchPageContentsInvalidURL(t *testing.T) {
	// Arrange
	client := fetcher.New(5 * time.Second)

	// Act
	url := "invalid-url"
	_, err := client.FetchPageContents(url)

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGrabLinksFromPageContents(t *testing.T) {
	// Arrange
	client := fetcher.New(5 * time.Second)
	url := "https://en.wikipedia.org/wiki/Initial_campaign_of_the_Breton_Civil_War"

	// Act
	links, err := client.GrabAllLinks(url)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Assert
	if len(links) == 0 {
		t.Fatalf("expected non-empty links, got empty slice")
	}
}
