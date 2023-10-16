package main

import (
	"net"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Url struct {
	Uri          string
	Active       bool
	responseTime time.Duration
	LastCalled   string
}

func main() {
	ticker := time.NewTicker(5 * time.Second)
	func() {
		for {
			select {
			case <-ticker.C:
				process()
			}
		}
	}()
}

func process() {
	var wg sync.WaitGroup
	mainCh := make(chan Url, 2)

	urlsToMonitor := []Url{
		Url{Uri: "google.com:80"},
		Url{Uri: "twitter.com:80"},
		Url{Uri: "abc.com:80"},
		Url{Uri: "http://ldkasldkas.com/"},
	}

	for _, v := range urlsToMonitor {
		wg.Add(1)
		go callUri(v.Uri, &wg, mainCh)
	}

	go func() {
		wg.Wait()
		close(mainCh)
	}()

	result := []Url{}
	for t := range mainCh {
		if t.Active != true {
			color.Red("URL: %s - Active: %t - LastCalled: %s\n", t.Uri, t.Active, t.LastCalled)
		} else {
			color.Green("URL: %s - Active: %t - ResponseTime: %s - LastCalled: %s\n", t.Uri, t.Active, t.responseTime, t.LastCalled)
		}
	}
	urlsToMonitor = result

}
func callUri(uri string, wg *sync.WaitGroup, mainCh chan<- Url) {
	defer wg.Done()

	conn, err := net.Dial("tcp", uri)

	if err != nil {
		mainCh <- Url{Uri: uri, responseTime: 0, Active: false, LastCalled: time.Now().String()}
		return
	}

	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))

	start := time.Now()
	oneByte := make([]byte, 1)
	_, err = conn.Read(oneByte)

	if err != nil {
		mainCh <- Url{Uri: uri, responseTime: time.Since(start), Active: false, LastCalled: time.Now().String()}
		return
	}

	mainCh <- Url{Uri: uri, responseTime: time.Since(start), Active: true, LastCalled: time.Now().String()}
}
