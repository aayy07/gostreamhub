package processing

import (
    "log"
    "gostreamhub/internal/storage"
)

// ProcessData handles data transformation, forwards results to the output channel, and saves processed data to the database
func ProcessData(input <-chan string, output chan<- string, dbStorage *storage.DBStorage) {
    for data := range input {
        log.Printf("Processing data: %s", data)

        // Process the data (this is just a placeholder for actual transformation logic)
        processedData := "Processed: " + data

        // Save processed data to the database
        if err := dbStorage.SaveProcessedData(data, processedData); err != nil {
            log.Printf("Failed to save data: %v", err)
        }

        // Send the processed data to the output channel
        output <- processedData
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
