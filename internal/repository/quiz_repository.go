package repository

import (
	"database/sql"
	"goserver/internal/domain"
)

type QuizRepository struct {
	db *sql.DB
}

func NewQuizRepository(db *sql.DB) *QuizRepository {
	return &QuizRepository{db}
}

func (s *QuizRepository) CreateQuiz(quiz domain.Quiz) (*domain.Quiz, error) {
	return nil, nil
}

func (s *QuizRepository) SubmitQuiz(resultQuiz domain.ResultQuiz) (*domain.ResultQuiz, error) {
	return nil, nil
}

func (s *QuizRepository) FindManyQuizzes() ([]domain.Quiz, error) {
	return nil, nil
}

func (s *QuizRepository) FindQuizById(id string) (*domain.Quiz, error) {
	return nil, nil
}