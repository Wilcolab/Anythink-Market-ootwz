package repository

import (
	"database/sql"
	"fmt"

	"github.com/Wilcolab/Anythink-Market-ootwz/models"
)

// QuestionRepository handles database operations for questions
type QuestionRepository struct {
	db *sql.DB
}

// NewQuestionRepository creates a new question repository
func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

// GetAll retrieves all questions from the database
func (r *QuestionRepository) GetAll() ([]models.Question, error) {
	query := `
		SELECT id, text, options, answer, category, created_at, updated_at 
		FROM questions 
		ORDER BY id ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query questions: %w", err)
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var q models.Question
		err := rows.Scan(
			&q.ID,
			&q.Text,
			&q.Options,
			&q.Answer,
			&q.Category,
			&q.CreatedAt,
			&q.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan question: %w", err)
		}
		questions = append(questions, q)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over questions: %w", err)
	}

	return questions, nil
}

// GetByID retrieves a specific question by ID
func (r *QuestionRepository) GetByID(id int) (*models.Question, error) {
	query := `
		SELECT id, text, options, answer, category, created_at, updated_at 
		FROM questions 
		WHERE id = $1
	`

	var q models.Question
	err := r.db.QueryRow(query, id).Scan(
		&q.ID,
		&q.Text,
		&q.Options,
		&q.Answer,
		&q.Category,
		&q.CreatedAt,
		&q.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Question not found
		}
		return nil, fmt.Errorf("failed to get question by ID %d: %w", id, err)
	}

	return &q, nil
}

// GetByCategory retrieves questions filtered by category
func (r *QuestionRepository) GetByCategory(category string) ([]models.Question, error) {
	query := `
		SELECT id, text, options, answer, category, created_at, updated_at 
		FROM questions 
		WHERE category = $1
		ORDER BY id ASC
	`

	rows, err := r.db.Query(query, category)
	if err != nil {
		return nil, fmt.Errorf("failed to query questions by category: %w", err)
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var q models.Question
		err := rows.Scan(
			&q.ID,
			&q.Text,
			&q.Options,
			&q.Answer,
			&q.Category,
			&q.CreatedAt,
			&q.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan question: %w", err)
		}
		questions = append(questions, q)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over questions: %w", err)
	}

	return questions, nil
}

// Create adds a new question to the database
func (r *QuestionRepository) Create(q *models.Question) error {
	query := `
		INSERT INTO questions (text, options, answer, category)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(
		query,
		q.Text,
		q.Options,
		q.Answer,
		q.Category,
	).Scan(&q.ID, &q.CreatedAt, &q.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create question: %w", err)
	}

	return nil
}

// Update modifies an existing question
func (r *QuestionRepository) Update(q *models.Question) error {
	query := `
		UPDATE questions 
		SET text = $2, options = $3, answer = $4, category = $5, updated_at = NOW()
		WHERE id = $1
		RETURNING updated_at
	`

	err := r.db.QueryRow(
		query,
		q.ID,
		q.Text,
		q.Options,
		q.Answer,
		q.Category,
	).Scan(&q.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("question with ID %d not found", q.ID)
		}
		return fmt.Errorf("failed to update question: %w", err)
	}

	return nil
}

// Delete removes a question from the database
func (r *QuestionRepository) Delete(id int) error {
	query := `DELETE FROM questions WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete question: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("question with ID %d not found", id)
	}

	return nil
}
