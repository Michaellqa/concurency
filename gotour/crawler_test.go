package gotour

import "testing"

func TestCrawl(t *testing.T) {
	Crawl("https://golang.org/", 4, fetcher)
}
