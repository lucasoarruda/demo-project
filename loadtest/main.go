package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	// Parse command-line arguments
	url := flag.String("url", "https://demo-project.example.com/", "The URL to make requests to")
	concurrency := flag.Int("concurrency", 10, "Number of concurrent requests to make")
	numRequests := flag.Int("num-requests", 10000, "Number of requests to make per goroutine")
	flag.Parse()

	// Print usage if no arguments are provided
	if flag.NFlag() == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	// WaitGroup to synchronize goroutines
	wg := sync.WaitGroup{}
	wg.Add(*concurrency)

	// Channel to receive response times
	responses := make(chan time.Duration, *concurrency**numRequests)

	// Spawn goroutines to make requests
	for i := 0; i < *concurrency; i++ {
		go func() {
			for j := 0; j < *numRequests; j++ {
				start := time.Now()
				resp, err := http.Get(*url)
				if err != nil {
					fmt.Println("Error making request:", err)
					continue
				}
				resp.Body.Close()
				elapsed := time.Since(start)
				responses <- elapsed
			}
			wg.Done()
		}()
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(responses)
	}()

	// Collect and print response times
	var total time.Duration
	var count int
	for r := range responses {
		total += r
		count++
		if count%1000 == 0 {
			log.Printf("Response time at count %d: %v\n", count, r)
		}
	}

	// Print summary statistics
	log.Println("--- Results ---")
	log.Printf("Requests made: %d\n", count)
	log.Printf("Total time taken: %v\n", total)
	log.Printf("Average response time: %v\n", total/time.Duration(count))
}
