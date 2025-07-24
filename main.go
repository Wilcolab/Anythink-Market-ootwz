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

// quizSubmitHandler handles POST /api/quiz/submit
func (app *App) quizSubmitHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		sendErrorResponse(w, "Method not allowed", "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request body
	var submission models.QuizSubmission
	if err := json.NewDecoder(r.Body).Decode(&submission); err != nil {
		sendErrorResponse(w, "Invalid JSON", "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate submission
	if len(submission.Answers) == 0 {
		sendErrorResponse(w, "Invalid submission", "At least one answer is required", http.StatusBadRequest)
		return
	}

	// Get all questions for validation and scoring
	questions, err := app.questionRepo.GetAll()
	if err != nil {
		log.Printf("Error fetching questions for scoring: %v", err)
		sendErrorResponse(w, "Database error", "Failed to fetch questions", http.StatusInternalServerError)
		return
	}

	// Create a map for quick question lookup
	questionMap := make(map[int]models.Question)
	for _, q := range questions {
		questionMap[q.ID] = q
	}

	// Validate answers and calculate score
	var results []models.AnswerResult
	correctAnswers := 0
	totalQuestions := len(submission.Answers)

	// Track answered question IDs to prevent duplicates
	answeredQuestions := make(map[int]bool)

	for _, answer := range submission.Answers {
		// Check for duplicate question IDs
		if answeredQuestions[answer.QuestionID] {
			sendErrorResponse(w, "Invalid submission", fmt.Sprintf("Duplicate answer for question ID %d", answer.QuestionID), http.StatusBadRequest)
			return
		}
		answeredQuestions[answer.QuestionID] = true

		// Find the question
		question, exists := questionMap[answer.QuestionID]
		if !exists {
			sendErrorResponse(w, "Invalid question", fmt.Sprintf("Question with ID %d does not exist", answer.QuestionID), http.StatusBadRequest)
			return
		}

		// Validate answer range
		if answer.Answer < 0 || answer.Answer >= len(question.Options) {
			sendErrorResponse(w, "Invalid answer", fmt.Sprintf("Answer %d is out of range for question %d", answer.Answer, answer.QuestionID), http.StatusBadRequest)
			return
		}

		// Check if answer is correct
		isCorrect := answer.Answer == question.Answer
		if isCorrect {
			correctAnswers++
		}

		// Add to results
		results = append(results, models.AnswerResult{
			QuestionID:    answer.QuestionID,
			UserAnswer:    answer.Answer,
			CorrectAnswer: question.Answer,
			IsCorrect:     isCorrect,
			Question:      question.Text,
		})
	}

	// Calculate percentage and determine pass/fail
	percentage := float64(correctAnswers) / float64(totalQuestions) * 100
	passed := percentage >= 60.0 // Pass threshold is 60%

	// Create response
	response := models.QuizResponse{
		Score:      correctAnswers,
		Total:      totalQuestions,
		Percentage: percentage,
		Passed:     passed,
		Results:    results,
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		sendErrorResponse(w, "Encoding failed", "Failed to encode quiz response", http.StatusInternalServerError)
		return
	}
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
	http.HandleFunc("/api/quiz/submit", app.quizSubmitHandler)

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
	fmt.Println("  POST /api/quiz/submit - Submit quiz answers and get score")
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
