package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/conceptcodes/webcrawler-go/internal/parser"
)

type Client struct {
	client     *http.Client
	timeout    time.Duration
	linkParser *parser.LinkParser
}

func New(timeout time.Duration) *Client {
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
		timeout:    timeout,
		linkParser: parser.NewLinkParser(),
	}
}

func (c *Client) GrabAllLinks(url string) ([]string, error) {
	content, err := c.FetchPageContents(url)
	if err != nil {
		return nil, err
	}

	links, err := c.linkParser.ParseLinks(content)
	if err != nil {
		return nil, err
	}

	return links, nil
}

func (c *Client) FetchPageContents(url string) (string, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch page: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
