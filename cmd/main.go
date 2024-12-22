package main

import (
    "fmt"
    "log"
    "net/http"
    "gostreamhub/internal/ingestion"
    "gostreamhub/internal/processing"
    "gostreamhub/internal/storage"
)

func main() {
    // Database configuration from environment or config file
    dbHost := "aayy07-db-aayy07.l.aivencloud.com"
    dbPort := "26366"
    dbUser := "avnadmin"
    dbPassword := "AVNS__knE3kZdUowOffhwGtn"
    dbName := "defaultdb"
    caCertPath := "C:\\Go-Workspace\\src\\ca.pem" // Path to your CA file

    // Initialize database connection
    dbStorage, err := storage.NewDBStorage(dbHost, dbPort, dbUser, dbPassword, dbName, caCertPath)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Create channels for communication
    inputChannel := make(chan string)
    outputChannel := make(chan string)

    // Start the processing pipeline
    go processing.ProcessData(inputChannel, outputChannel, dbStorage)

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
