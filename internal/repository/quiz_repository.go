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

func (s *QuizRepository) FindManyQuizzes() ([]domain.Quiz, error) {
	return nil, nil
}
