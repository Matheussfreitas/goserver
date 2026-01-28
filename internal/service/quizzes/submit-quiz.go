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

func (s *SubmitQuizService) SubmitQuiz(resultQuiz domain.ResultQuiz, userId string) (*domain.ResultQuiz, error) {
	quiz, err := s.repo.FindQuizById(resultQuiz.QuizID, userId)
	if err != nil {
		return nil, err
	}
	if quiz == nil {
		return nil, sql.ErrNoRows
	}

	score := 0
	questionsMap := make(map[string]domain.Question)
	for _, q := range quiz.Questions {
		questionsMap[q.ID] = q
	}

	for i, ans := range resultQuiz.Answers {
		if q, ok := questionsMap[ans.QuestionID]; ok {
			isCorrect := q.CorrectAnswer == ans.UserChoice
			resultQuiz.Answers[i].IsCorrect = isCorrect
			if isCorrect {
				score++
			}
		}
	}

	resultQuiz.Score = score
	resultQuiz.TotalQuestions = len(quiz.Questions)

	return s.repo.SubmitQuiz(resultQuiz)
}
