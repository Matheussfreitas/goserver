package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"goserver/internal/domain"
)

type QuizRepository struct {
	db *sql.DB
}

func NewQuizRepository(db *sql.DB) *QuizRepository {
	return &QuizRepository{db}
}

func (s *QuizRepository) CreateQuiz(userId string, quiz domain.Quiz) (*domain.Quiz, error) {
	query := `
	INSERT INTO quizzes (
		user_id,
		title,
		content,
		difficulty,
		number_questions,
		created_at
	) VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, created_at
	`

	var row *sql.Row

	row = s.db.QueryRowContext(context.Background(), query, userId, quiz.Title, quiz.Content, quiz.Difficulty, quiz.NumberQuestions, quiz.CreatedAt)

	if err := row.Scan(&quiz.ID, &quiz.CreatedAt); err != nil {
		return nil, err
	}

	for i, q := range quiz.Questions {
		answersJSON, err := json.Marshal(q.Answers)
		if err != nil {
			return nil, err
		}

		_, err = s.db.ExecContext(context.Background(), `
			INSERT INTO questions (quiz_id, statement, answers, correct_answer, explanation)
			VALUES ($1, $2, $3, $4, $5)
		`, quiz.ID, q.Statement, answersJSON, q.CorrectAnswer, q.Explanation)

		if err != nil {
			return nil, err
		}
		quiz.Questions[i].QuizID = quiz.ID
	}

	return &quiz, nil
}

func (s *QuizRepository) SubmitQuiz(resultQuiz domain.ResultQuiz) (*domain.ResultQuiz, error) {
	return nil, nil
}

func (s *QuizRepository) FindManyQuizzes(userId string) ([]domain.Quiz, error) {
	query := `SELECT * FROM quizzes WHERE user_id = $1`

	rows, err := s.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizzes []domain.Quiz
	for rows.Next() {
		var quiz domain.Quiz
		if err := rows.Scan(&quiz.ID, &quiz.UserID, &quiz.Title, &quiz.Content, &quiz.Difficulty, &quiz.NumberQuestions, &quiz.CreatedAt); err != nil {
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}
	return quizzes, nil
}

func (s *QuizRepository) FindQuizById(id, userId string) (*domain.Quiz, error) {
	query := `SELECT * FROM quizzes WHERE id = $1 AND user_id = $2`

	rows, err := s.db.Query(query, id, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quiz domain.Quiz
	if rows.Next() {
		if err := rows.Scan(&quiz.ID, &quiz.UserID, &quiz.Title, &quiz.Content, &quiz.Difficulty, &quiz.NumberQuestions, &quiz.CreatedAt); err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}

	questionsQuery := `SELECT id, quiz_id, statement, answers, correct_answer, explanation FROM questions WHERE quiz_id = $1`
	questionRows, err := s.db.Query(questionsQuery, quiz.ID)
	if err != nil {
		return nil, err
	}
	defer questionRows.Close()

	for questionRows.Next() {
		var q domain.Question
		var answersJSON []byte
		if err := questionRows.Scan(&q.ID, &q.QuizID, &q.Statement, &answersJSON, &q.CorrectAnswer, &q.Explanation); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(answersJSON, &q.Answers); err != nil {
			return nil, err
		}
		quiz.Questions = append(quiz.Questions, q)
	}

	return &quiz, nil
}
