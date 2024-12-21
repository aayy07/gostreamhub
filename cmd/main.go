package main

import (
	"fmt"
	"gostreamhub/internal/ingestion"
	"gostreamhub/internal/processing"
	"log"
	"net/http"
)

func main() {
	// Create channels for communication
	inputChannel := make(chan string)
	outputChannel := make(chan string)

	// Start the processing pipeline
	go processing.ProcessData(inputChannel, outputChannel)

	// Log processed data for demonstration
	go func() {
		for processed := range outputChannel {
			log.Printf("Processed Data Ready for Output: %s", processed)
		}
	}()

	// Set up the HTTP ingestion endpoint
	http.HandleFunc("/ingest", ingestion.HTTPIngest(inputChannel))

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


// Logic
// The main.go file ties everything together. It initializes the application, sets up channels, starts the processing pipeline, and serves the HTTP endpoint.

// Explanation
// Channels:
// inputChannel: Receives raw data from ingestion.
// outputChannel: Collects processed data from the pipeline.
// Processing Pipeline: Runs in a goroutine to continuously process data.
// HTTP Server: Serves the /ingest endpoint for receiving data.