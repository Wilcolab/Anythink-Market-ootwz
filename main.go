package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Wilcolab/Anythink-Market-ootwz/config"
	"github.com/Wilcolab/Anythink-Market-ootwz/database"
	"github.com/Wilcolab/Anythink-Market-ootwz/models"
	"github.com/Wilcolab/Anythink-Market-ootwz/repository"
	_ "github.com/lib/pq"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// App holds the application dependencies
type App struct {
	questionRepo *repository.QuestionRepository
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
func (app *App) questionsHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET method
	if r.Method != http.MethodGet {
		sendErrorResponse(w, "Method not allowed", "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get optional category filter from query parameters
	category := r.URL.Query().Get("category")

	var questions []models.Question
	var err error

	if category != "" {
		questions, err = app.questionRepo.GetByCategory(category)
	} else {
		questions, err = app.questionRepo.GetAll()
	}

	if err != nil {
		log.Printf("Error fetching questions: %v", err)
		sendErrorResponse(w, "Database error", "Failed to fetch questions", http.StatusInternalServerError)
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
func (app *App) questionByIDHandler(w http.ResponseWriter, r *http.Request) {
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

	// Get question from database
	question, err := app.questionRepo.GetByID(id)
	if err != nil {
		log.Printf("Error fetching question by ID %d: %v", id, err)
		sendErrorResponse(w, "Database error", "Failed to fetch question", http.StatusInternalServerError)
		return
	}

	// Check if question was found
	if question == nil {
		sendErrorResponse(w, "Question not found", fmt.Sprintf("Question with ID %d does not exist", id), http.StatusNotFound)
		return
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send JSON response
	if err := json.NewEncoder(w).Encode(question); err != nil {
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
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(db, "migrations"); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize repositories
	questionRepo := repository.NewQuestionRepository(db)

	// Initialize app with dependencies
	app := &App{
		questionRepo: questionRepo,
	}

	// Health check endpoint
	http.HandleFunc("/health", healthCheckHandler)

	// API endpoints
	http.HandleFunc("/api/questions", app.questionsHandler)
	http.HandleFunc("/api/questions/", app.questionByIDHandler)

	// Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go Quiz API! 🐹")
	})

	fmt.Printf("🚀 Server starting on port %s...\n", cfg.Port)
	fmt.Println("📋 Available endpoints:")
	fmt.Println("  GET /health - Health check")
	fmt.Println("  GET /api/questions - List all questions")
	fmt.Println("  GET /api/questions?category={category} - Filter questions by category")
	fmt.Println("  GET /api/questions/{id} - Get question by ID")
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
