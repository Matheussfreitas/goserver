package quizzes

import (
	"database/sql"
	"goserver/internal/domain"
	"goserver/internal/repository"
)

type SubmitQuizService struct {
	repo *repository.QuizRepository
	db   *sql.DB
}

func NewSubmitQuizService(repo *repository.QuizRepository, db *sql.DB) *SubmitQuizService {
	return &SubmitQuizService{
		repo: repo,
		db:   db,
	}
}

func (s *SubmitQuizService) SubmitQuiz(resultQuiz domain.ResultQuiz) (*domain.ResultQuiz, error) {
	return s.repo.SubmitQuiz(resultQuiz)
}
