package domain

import "time"

type Quiz struct {
	ID              string         `json:"id"`
	UserID          string         `json:"user_id"`
	Title           string         `json:"title"`
	Content         string         `json:"content"`
	Difficulty      QuizDifficulty `json:"difficulty"`
	NumberQuestions int            `json:"number_questions"`
	Questions       []Question     `json:"questions,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
}

type QuizDifficulty string

const (
	Easy   QuizDifficulty = "easy"
	Medium QuizDifficulty = "medium"
	Hard   QuizDifficulty = "hard"
)
