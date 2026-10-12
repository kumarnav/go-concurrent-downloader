package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Struct to hold the result of each URL check
type Result struct {
	URL     string
	Status  string
	Code    int
	Latency time.Duration
}

func checkURL(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()

	start := time.Now()
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	latency := time.Since(start)

	if err != nil {
		results <- Result{URL: url, Status: "FAILED", Code: 0, Latency: latency}
		return
	}
	defer resp.Body.Close()

	results <- Result{URL: url, Status: "SUCCESS", Code: resp.StatusCode, Latency: latency}
}

func main() {
	// List of sample URLs to check concurrently
	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://httpbin.org/delay/1",
		"https://api.github.com",
		"https://google.com",
		"https://stackoverflow.com",
	}

	fmt.Printf("[*] Starting concurrent check for %d URLs...\n\n", len(urls))

	startTime := time.Now()

	var wg sync.WaitGroup
	results := make(chan Result, len(urls))

	// Launch a goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go checkURL(url, &wg, results)
	}

	// Wait for all goroutines to finish in the background
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results as they arrive concurrently
	for res := range results {
		if res.Status == "SUCCESS" {
			fmt.Printf("[+] [%s] Code: %d | Latency: %v | URL: %s\n", res.Status, res.Code, res.Latency, res.URL)
		} else {
			fmt.Printf("[-] [%s] Latency: %v | URL: %s\n", res.Status, res.Latency, res.URL)
		}
	}

	fmt.Printf("\n[+] All tasks completed in %v\n", time.Since(startTime))
}