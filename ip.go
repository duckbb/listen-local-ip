package listen_local_ip

import (
	"encoding/json"
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

type info struct {
	Cip   string `json:"cip"`
	Cid   string `json:"cid"`
	Cname string `json:"cname"`
}

func (s *Service) Get() {
	c := colly.NewCollector()
	c.OnResponse(func(resp *colly.Response) {
		if len(resp.Body) == 0 {
			return
		}
		t := &info{}
		ipRegexp := regexp.MustCompile("{.*}")
		prevIp := ipRegexp.FindString(string(resp.Body))
		err := json.Unmarshal([]byte(prevIp), t)
		if err != nil {
			log.Println("data unmarshal err:", err)
		}
		if t.Cip != "" {
			Write(t.Cip)
		}
	})
	c.Visit("http://pv.sohu.com/cityjson?ie=utf-8")
}
