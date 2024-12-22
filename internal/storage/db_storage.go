package storage

import (
    "database/sql"
    "fmt"
    "log"

    // Import the pq package to register the PostgreSQL driver
    _ "github.com/lib/pq" 
)

// DBStorage struct to manage database connections
type DBStorage struct {
    db *sql.DB
}

// NewDBStorage initializes a new DBStorage instance
func NewDBStorage(host, port, user, password, dbName, caCertPath string) (*DBStorage, error) {
    // Build the DSN with sslmode=require for secure connections
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=require sslrootcert=%s",
        host, port, user, password, dbName, caCertPath,
    )

    // Open the database connection
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to the database: %v", err)
    }

    // Verify the connection
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping the database: %v", err)
    }

    log.Println("Successfully connected to the database!")
    return &DBStorage{db: db}, nil
}

// Add your SaveProcessedData and other functions for DB interaction here...


// SaveProcessedData saves the processed data into the database
func (ds *DBStorage) SaveProcessedData(originalData, processedData string) error {
    query := "INSERT INTO processed_data (original_data, processed_data) VALUES ($1, $2)"
    _, err := ds.db.Exec(query, originalData, processedData)
    if err != nil {
        log.Printf("Failed to save data (original: %s, processed: %s): %v", originalData, processedData, err)
        return fmt.Errorf("failed to save processed data: %v", err)
    }

    return nil
}
