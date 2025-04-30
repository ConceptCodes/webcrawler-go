package parser_test

import (
	"testing"

	"github.com/conceptcodes/webcrawler-go/internal/parser"
)

func TestParseLinks(t *testing.T) {
	// Arrange
	content := `<html><body><a href="http://example.com">Example</a></body></html>`
	lp := parser.NewLinkParser()

	// Act
	links, err := lp.ParseLinks(content)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(links) != 1 {
		t.Fatalf("expected 1 link, got %d", len(links))
	}
	if links[0] != "http://example.com" {
		t.Fatalf("expected http://example.com, got %s", links[0])
	}
}
