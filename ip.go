package listen_local_ip

import (
	"github.com/gocolly/colly"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Get() {

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		if e.Text != "" {
			Write(e.Text)
		}
	})

	c.Visit("http://202020.ip138.com/")
}
