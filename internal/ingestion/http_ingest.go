package ingestion

import (
	"fmt"
	"net/http"
)

// HTTPIngest receives data via HTTP and forwards it to the input channel.
func HTTPIngest(input chan<- string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Query().Get("data")
		if data == "" {
			http.Error(w, "No data provided", http.StatusBadRequest)
			return
		}

		// Forward data to the processing pipeline
		input <- data
		fmt.Fprintf(w, "Received: %s", data)
	}
}


// Logic:
// Ingest data via an HTTP endpoint. This simulates receiving data from an external source and forwarding it to the processing pipeline.

// Explanation:
// HTTP Handler: Reads data from the query string.
// Validation: Returns a 400 error if no data is provided.
// Channel Communication: Sends data to the input channel for processing.