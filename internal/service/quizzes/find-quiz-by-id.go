package quizzes

import (
	 "database/sql"
	 "goserver/internal/repository"
	 "goserver/internal/domain"
)

type FindQuizByIdService struct {
	repo *repository.QuizRepository
	db   *sql.DB
}

func NewFindQuizByIdService(repo *repository.QuizRepository, db *sql.DB) *FindQuizByIdService {
	return &FindQuizByIdService{
		repo: repo,
		db:   db,
	}
}

func (s *FindQuizByIdService) FindQuizById(id, userId string) (*domain.Quiz, error) {
	return s.repo.FindQuizById(id, userId)
}


