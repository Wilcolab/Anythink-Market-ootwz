package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Question represents a quiz question
type Question struct {
	ID       int      `json:"id"`
	Text     string   `json:"text"`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"`
	Category string   `json:"category"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// In-memory questions data
var questions = []Question{
	{
		ID:       1,
		Text:     "What is the capital of France?",
		Options:  []string{"London", "Berlin", "Paris", "Madrid"},
		Answer:   2,
		Category: "Geography",
	},
	{
		ID:       2,
		Text:     "Which programming language is known for its simplicity and efficiency?",
		Options:  []string{"Java", "Go", "C++", "Python"},
		Answer:   1,
		Category: "Programming",
	},
	{
		ID:       3,
		Text:     "What is 2 + 2?",
		Options:  []string{"3", "4", "5", "6"},
		Answer:   1,
		Category: "Math",
	},
	{
		ID:       4,
		Text:     "Who wrote 'Romeo and Juliet'?",
		Options:  []string{"Charles Dickens", "William Shakespeare", "Jane Austen", "Mark Twain"},
		Answer:   1,
		Category: "Literature",
	},
}

// healthCheckHandler handles the /health endpoint
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Create health response
	response := HealthResponse{
		Status:  "healthy",
		Message: "Server is running successfully",
	}

	// Encode and send JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

// questionsHandler handles GET /api/questions
func questionsHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET method
	if r.Method != http.MethodGet {
		sendErrorResponse(w, "Method not allowed", "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send JSON response
	if err := json.NewEncoder(w).Encode(questions); err != nil {
		sendErrorResponse(w, "Encoding failed", "Failed to encode questions response", http.StatusInternalServerError)
		return
	}
}

// questionByIDHandler handles GET /api/questions/{id}
func questionByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET method
	if r.Method != http.MethodGet {
		sendErrorResponse(w, "Method not allowed", "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/questions/")
	if path == "" {
		sendErrorResponse(w, "Invalid request", "Question ID is required", http.StatusBadRequest)
		return
	}

	// Parse ID
	id, err := strconv.Atoi(path)
	if err != nil {
		sendErrorResponse(w, "Invalid ID", "Question ID must be a valid number", http.StatusBadRequest)
		return
	}

	// Find question by ID
	var foundQuestion *Question
	for i := range questions {
		if questions[i].ID == id {
			foundQuestion = &questions[i]
			break
		}
	}

	// Check if question was found
	if foundQuestion == nil {
		sendErrorResponse(w, "Question not found", fmt.Sprintf("Question with ID %d does not exist", id), http.StatusNotFound)
		return
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send JSON response
	if err := json.NewEncoder(w).Encode(foundQuestion); err != nil {
		sendErrorResponse(w, "Encoding failed", "Failed to encode question response", http.StatusInternalServerError)
		return
	}
}

// sendErrorResponse sends a JSON error response
func sendErrorResponse(w http.ResponseWriter, error, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := ErrorResponse{
		Error:   error,
		Message: message,
	}

	json.NewEncoder(w).Encode(errorResponse)
}

func main() {
	// Connect to database
	db := connectDB()
	defer db.Close()

	// Health check endpoint
	http.HandleFunc("/health", healthCheckHandler)

	// API endpoints
	http.HandleFunc("/api/questions", questionsHandler)
	http.HandleFunc("/api/questions/", questionByIDHandler)

	// Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go API! 🐹")
	})

	fmt.Println("🚀 Server starting on port 8080...")
	fmt.Println("📋 Available endpoints:")
	fmt.Println("  GET /health - Health check")
	fmt.Println("  GET /api/questions - List all questions")
	fmt.Println("  GET /api/questions/{id} - Get question by ID")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connectDB() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("✅ Connected to database")
	return db
}
