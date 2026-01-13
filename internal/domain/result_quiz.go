package domain

import "time"

type ResultQuiz struct {
	ID             string    `json:"id"`
	QuizID         string    `json:"quiz_id"`
	UserID         string    `json:"user_id"`
	Score          int       `json:"score"`
	TotalQuestions int       `json:"total_questions"`
	Answers        []UserAnswer  `json:"answers"`
	CompletedAt    time.Time `json:"completed_at"`
}

type UserAnswer struct {
	QuestionID string `json:"question_id"` // UUID
	UserChoice int    `json:"user_choice"`
	IsCorrect  bool   `json:"is_correct"`
}
