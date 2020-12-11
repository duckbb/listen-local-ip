package main

import (
	"flag"
	"sync"
	"time"

	Ipserver "github.com/duckbb/listen-local-ip"
)

var etime int

func main() {
	flag.IntVar(&etime, "t", 10, "monitoy interval ")
	flag.Parse()
	Ipserver.Init()
	var wg sync.WaitGroup
	ticker := time.NewTicker(time.Second * time.Duration(etime))
	for {
		Ipserver.Log.Println("start-time:", time.Now())
		<-ticker.C
		wg.Add(1)
		go func(t *time.Ticker) {
			defer wg.Done()
			defer func() {
				if err := recover(); err != nil {
					Ipserver.Log.Println("server fail! err:", err)
					return
				}
			}()
			server := Ipserver.NewService()
			server.Get()

		}(ticker)
		Ipserver.Log.Println("end-time:", time.Now())
	}
	wg.Wait()
}
