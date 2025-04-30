package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client

	timeout time.Duration
}

func New(timeout time.Duration) *Client {
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
		timeout: timeout,
	}
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
