package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Question represents a quiz question in the database
type Question struct {
	ID        int       `json:"id" db:"id"`
	Text      string    `json:"text" db:"text"`
	Options   Options   `json:"options" db:"options"`
	Answer    int       `json:"answer" db:"answer"`
	Category  string    `json:"category" db:"category"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Options represents the question options stored as JSONB in PostgreSQL
type Options []string

// Value implements the driver.Valuer interface for database storage
func (o Options) Value() (driver.Value, error) {
	return json.Marshal(o)
}

// Scan implements the sql.Scanner interface for database retrieval
func (o *Options) Scan(value interface{}) error {
	if value == nil {
		*o = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, o)
	case string:
		return json.Unmarshal([]byte(v), o)
	default:
		return errors.New("cannot scan into Options")
	}
}

// QuizResult represents a user's quiz attempt result
type QuizResult struct {
	ID          int       `json:"id" db:"id"`
	UserID      string    `json:"user_id" db:"user_id"`
	QuestionID  int       `json:"question_id" db:"question_id"`
	UserAnswer  int       `json:"user_answer" db:"user_answer"`
	IsCorrect   bool      `json:"is_correct" db:"is_correct"`
	SubmittedAt time.Time `json:"submitted_at" db:"submitted_at"`
}

// QuizSession represents a complete quiz session
type QuizSession struct {
	ID             int        `json:"id" db:"id"`
	UserID         string     `json:"user_id" db:"user_id"`
	Score          int        `json:"score" db:"score"`
	TotalQuestions int        `json:"total_questions" db:"total_questions"`
	StartedAt      time.Time  `json:"started_at" db:"started_at"`
	CompletedAt    *time.Time `json:"completed_at,omitempty" db:"completed_at"`
}

// QuizAnswer represents a single answer in a quiz submission
type QuizAnswer struct {
	QuestionID int `json:"questionId" validate:"required"`
	Answer     int `json:"answer" validate:"required,min=0"`
}

// QuizSubmission represents the request body for quiz submission
type QuizSubmission struct {
	Answers []QuizAnswer `json:"answers" validate:"required,min=1"`
}

// QuizResponse represents the response after quiz submission
type QuizResponse struct {
	Score      int            `json:"score"`
	Total      int            `json:"total"`
	Percentage float64        `json:"percentage"`
	Passed     bool           `json:"passed"`
	Results    []AnswerResult `json:"results,omitempty"`
}

// AnswerResult represents the result for a single question
type AnswerResult struct {
	QuestionID    int    `json:"questionId"`
	UserAnswer    int    `json:"userAnswer"`
	CorrectAnswer int    `json:"correctAnswer"`
	IsCorrect     bool   `json:"isCorrect"`
	Question      string `json:"question,omitempty"`
}
