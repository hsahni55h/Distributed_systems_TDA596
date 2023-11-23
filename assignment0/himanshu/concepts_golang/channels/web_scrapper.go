package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// fetchURL fetches a URL, processes the response, and sends the result or error through channels
func fetchURL(url string, results chan<- string, errors chan<- error) {
	// Make an HTTP GET request to the specified URL
	resp, err := http.Get(url)
	if err != nil {
		// If there's an error during the HTTP request, send the error to the errors channel and return
		errors <- err
		return
	}
	defer resp.Body.Close() // Ensure the response body is closed when the function returns

	// Read the body of the response using io.ReadAll
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// If there's an error reading the response body, send the error to the errors channel and return
		errors <- err
		return
	}

	// Simulate some processing time
	time.Sleep(2 * time.Second)

	// Send the result to the results channel along with the URL
	results <- fmt.Sprintf("%s: %d bytes", url, len(body))
}

func main() {
	// List of URLs to scrape
	urls := []string{"https://www.example.com", "https://www.example.org", "https://www.example.net", "https://nonexistent"}

	// Create channels for results and errors
	results := make(chan string)
	errors := make(chan error)

	// Start a goroutine for each URL
	for _, url := range urls {
		go fetchURL(url, results, errors)
	}

	// Wait for all goroutines to finish
	for range urls {
		select {
		case result := <-results:
			// Case 1: Received a result from one of the goroutines
			fmt.Println(result)
		case err := <-errors:
			// Case 2: Received an error from one of the goroutines
			fmt.Println("Error:", err)
		}
	}
}
