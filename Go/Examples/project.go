/*
Project Activity: Concurrent Web Scraper using Goroutines

                        Objective:
Create a concurrent web scraper in Go that fetches content from multiple URLs simultaneously.

This project will introduce the concept of using goroutines to handle multiple I/O-bound tasks (HTTP requests) concurrently, improving the efficiency of the scraping process.

                        Scenario:
You are tasked with building a web scraper that can fetch the titles of multiple websites concurrently.

The scraper should also handle possible errors such as timeouts, connection issues, and invalid URLs.

                        Requirements:
Concurrency: Use goroutines to scrape multiple URLs concurrently.

Synchronization: Use a sync.WaitGroup to ensure all goroutines complete before the program exits.

Error Handling: Handle possible errors during HTTP requests, such as connection timeouts, invalid URLs, and unavailable websites.

Channel Communication: Use channels to collect results from each goroutine and display them at the end.

Rate Limiting: (Optional) Implement a rate limiter to prevent overwhelming the target websites with too many concurrent requests.


            Instructions for Students:

1. Concurrency with Goroutines: Each URL is scraped concurrently by spawning a separate goroutine for each URL.

2. Synchronization: The sync.WaitGroup ensures that the main program waits until all goroutines complete before exiting.

3. Channel Communication: A channel is used to collect and print the results from each goroutine in a synchronized manner.

4. Error Handling: Proper error handling is implemented to manage HTTP errors (e.g., timeouts, invalid URLs).
*/

package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

// fetchTitle fetches the page content and returns the title of the web page
func fetchTitle(url string) (string, error) {
	// Make an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("Failed to fetch %s: %s", url, resp.Status))
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Find and return the title tag
	bodyStr := string(body)
	startIdx := strings.Index(bodyStr, "<title>")
	endIdx := strings.Index(bodyStr, "</title>")
	if startIdx == -1 || endIdx == -1 {
		return "", errors.New("No title found in page")
	}

	title := bodyStr[startIdx+len("<title>") : endIdx]
	return strings.TrimSpace(title), nil
}

// worker function fetches the title of a URL and sends the result through a channel
func worker(url string, wg *sync.WaitGroup, resultChan chan<- string) {
	defer wg.Done()

	// Fetch the title
	title, err := fetchTitle(url)
	if err != nil {
		resultChan <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}

	resultChan <- fmt.Sprintf("Title of %s: %s", url, title)
}

func main() {
	// List of URLs to scrape
	urls := []string{
		"https://golang.org",
		"https://www.github.com",
		"https://www.google.com",
		"https://news.ycombinator.com",
		"https://www.reddit.com",
	}

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a channel to collect results from goroutines
	resultChan := make(chan string, len(urls))

	// Start a goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go worker(url, &wg, resultChan)
	}

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Print the results from the channel
	for result := range resultChan {
		fmt.Println(result)
	}
}
