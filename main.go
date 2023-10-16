package main

import (
	"log"
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
	tickerInterval := 5 * time.Second
	ticker := time.NewTicker(tickerInterval)

	urlsToMonitor := []Url{
		{Uri: "google.com:80"},
		{Uri: "ldkasldkas.com:80"},
	}

	for range ticker.C {
		process(urlsToMonitor)
	}
}

func process(urls []Url) {
	var wg sync.WaitGroup
	mainCh := make(chan Url, 2)

	for _, url := range urls {
		wg.Add(1)
		go callUri(url.Uri, &wg, mainCh)
	}

	go func() {
		wg.Wait()
		close(mainCh)
	}()

	for urlResult := range mainCh {
		if urlResult.Active {
			color.Green("URL: %s - Active: %t - ResponseTime: %s - LastCalled: %s\n", urlResult.Uri, urlResult.Active, urlResult.responseTime, urlResult.LastCalled)
		} else {
			color.Red("URL: %s - Active: %t - LastCalled: %s\n", urlResult.Uri, urlResult.Active, urlResult.LastCalled)
		}
	}
}

func callUri(uri string, wg *sync.WaitGroup, mainCh chan<- Url) {
	defer wg.Done()

	conn, err := net.Dial("tcp", uri)
	lastCalled := time.Now().String()

	if err != nil {
		log.Printf("Failed to connect to URL: %s - Error: %v\n", uri, err)
		mainCh <- Url{Uri: uri, responseTime: 0, Active: false, LastCalled: lastCalled}
		return
	}

	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))

	start := time.Now()
	oneByte := make([]byte, 1)
	_, err = conn.Read(oneByte)

	if err != nil {
		log.Printf("Failed to read from URL: %s - Error: %v\n", uri, err)
		mainCh <- Url{Uri: uri, responseTime: time.Since(start), Active: false, LastCalled: lastCalled}
		return
	}

	mainCh <- Url{Uri: uri, responseTime: time.Since(start), Active: true, LastCalled: lastCalled}
}
