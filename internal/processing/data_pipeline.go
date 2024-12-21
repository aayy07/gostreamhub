package processing

import (
	"log"
	"strings"
)

// ProcessData handles data transformation and forwards the results to the output channel.
func ProcessData(input <-chan string, output chan<- string) {
	for msg := range input {
		// Example transformation: Convert the message to uppercase
		processed := strings.ToUpper(msg)
		log.Printf("Processing data: %s -> %s", msg, processed)
		output <- processed
	}
}


// Logic:
// The processing logic transforms raw data into a more usable format. Here, we'll simulate this by:

// Receiving raw data from an input channel.
// Transforming the data (e.g., converting text to uppercase).
// Sending the processed data to an output channel.

// Explanation:
// Input Channel: Receives data from the ingestion component.
// Transformation: Applies a simple transformation (convert text to uppercase).
// Output Channel: Forwards the processed data for further handling (e.g., storage or output).