package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// LogMessage represents the structure of messages from jarvis-agent-windows
type LogMessage struct {
	Timestamp   time.Time              `json:"timestamp"`
	Source      string                 `json:"source"`
	Type        string                 `json:"type"`
	Level       string                 `json:"level"`
	Message     string                 `json:"message"`
	Data        map[string]interface{} `json:"data,omitempty"`
	AgentID     string                 `json:"agent_id"`
	Environment string                 `json:"environment"`
}

func main() {
	// Create log file
	logFile, err := os.OpenFile("api_server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	// Configure logging to write to both file and console
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	// Define routes
	http.HandleFunc("/api/v1/logs", handleLogs)

	// Start server
	port := 8080
	log.Printf("Jarvis API Test Server starting on port %d...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func handleLogs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleLogPost(w, r)
	case http.MethodGet:
		handleLogGet(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleLogPost(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Parse the message
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Print the received data
	prettyJSON, _ := json.MarshalIndent(data, "", "  ")
	log.Printf("Received log message from agent %s\n", string(prettyJSON))

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Log received",
	})
}

func handleLogGet(w http.ResponseWriter, _ *http.Request) {
	// For now, just return a simple message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Log retrieval not implemented",
	})
}
