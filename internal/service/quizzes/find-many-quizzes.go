package quizzes

import (
	"database/sql"
	"goserver/internal/repository"
	"goserver/internal/domain"
)

type FindManyQuizzesService struct {
	db *sql.DB
	repo *repository.QuizRepository
}

func NewFindManyQuizzesService(repo *repository.QuizRepository, db *sql.DB) *FindManyQuizzesService {
	return &FindManyQuizzesService{
		repo: repo,
		db: db,
	}
}

func (s *FindManyQuizzesService) FindManyQuizzes() ([]domain.Quiz, error) {
	return s.repo.FindManyQuizzes()
}
