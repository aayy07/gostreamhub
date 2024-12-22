package ingestion

import (
	"fmt"
	"log"
	"net/http"
)

// HTTPIngest creates an HTTP handler that writes incoming data to the channel
func HTTPIngest(inputChannel chan<- string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Query().Get("data")
		if data == "" {
			http.Error(w, "Missing 'data' parameter", http.StatusBadRequest)
			log.Println("Error: Missing 'data' parameter in request")
			return
		}

		// Log the received data
		log.Printf("Received data: %s", data)

		// Send data to the processing pipeline
		select {
		case inputChannel <- data:
			fmt.Fprintln(w, "Data ingested successfully")
			log.Printf("Data sent to input channel: %s", data)
		default:
			http.Error(w, "Input channel is full, try again later", http.StatusServiceUnavailable)
			log.Println("Error: Input channel is full")
		}
	}
}

// Logic:
// Ingest data via an HTTP endpoint. This simulates receiving data from an external source and forwarding it to the processing pipeline.

// Explanation:
// HTTP Handler: Reads data from the query string.
// Validation: Returns a 400 error if no data is provided.
// Channel Communication: Sends data to the input channel for processing.
