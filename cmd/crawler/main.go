package crawler

import (
	"time"

	"github.com/conceptcodes/webcrawler-go/internal/scheduler"
	"github.com/conceptcodes/webcrawler-go/internal/fetcher"
)

type Crawler struct {
	BaseUrl string
	Queue   *scheduler.SimpleQueue
	client  *fetcher.Client
}

func New(baseUrl string, timeout time.Duration) *Crawler {
	return &Crawler{
		BaseUrl: baseUrl,
		Queue:   scheduler.New(),
		client:  fetcher.New(timeout),
	}
}

func (c *Crawler) Start() {
	links, err := c.client.GrabAllLinks(c.BaseUrl)
	if err != nil {
		println("Error fetching links:", err)
		return
	}
	for _, link := range links {
		c.AddUrl(link)
	}
	c.ProcessQueue()
}

func (c *Crawler) AddUrl(url string) {
	c.Queue.Enqueue(url)
	println("Added URL to queue:", url)
}

func (c *Crawler) ProcessQueue() {
	for !c.Queue.IsEmpty() {
		url := c.Queue.Dequeue()
		println("Processing URL:", url)
	}
}
