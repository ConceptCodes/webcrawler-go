package parser

import (
	"strings"

	"golang.org/x/net/html"
)

type LinkParser struct {
}

func NewLinkParser() *LinkParser {
	return &LinkParser{}
}

var f func(*html.Node)

func (lp *LinkParser) ParseLinks(content string) ([]string, error) {
	var links []string

	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return nil, err
	}
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links, nil
}
