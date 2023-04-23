package internal

import "context"

type Crawler interface {
	Crawl(ctx context.Context)
}

func NewTravelerCrawler() Crawler {
	return &travelerCrawler{}
}

type travelerCrawler struct {
}

func (t *travelerCrawler) Crawl(ctx context.Context) {

}
