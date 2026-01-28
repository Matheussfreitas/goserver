package handler

import (
	"database/sql"
	"encoding/json"
	"goserver/internal/domain"
	"goserver/internal/repository"
	"goserver/internal/service/quizzes"
	"net/http"
	"time"
)

type QuizHandler struct {
	findManyQuizzesService *quizzes.FindManyQuizzesService
	createQuizService      *quizzes.CreateQuizService
	findQuizByIdService    *quizzes.FindQuizByIdService
	submitQuizService      *quizzes.SubmitQuizService
}

func NewQuizHandler(db *sql.DB) *QuizHandler {
	repo := repository.NewQuizRepository(db)
	return &QuizHandler{
		findManyQuizzesService: quizzes.NewFindManyQuizzesService(repo, db),
		createQuizService:      quizzes.NewCreateQuizService(repo, db),
		findQuizByIdService:    quizzes.NewFindQuizByIdService(repo, db),
		submitQuizService:      quizzes.NewSubmitQuizService(repo, db),
	}
}

type CreateQuizRequest struct {
	Tema        string `json:"tema"`
	NumQuestoes int    `json:"numQuestoes"`
	Dificuldade string `json:"dificuldade"`
}

type SubmitQuizRequest struct {
	QuizID  string              `json:"quiz_id"`
	UserID  string              `json:"user_id"`
	Answers []domain.UserAnswer `json:"answers"`
}

func (h *QuizHandler) CreateQuiz(w http.ResponseWriter, r *http.Request) {
	var req CreateQuizRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	userId := r.Context().Value("user_id").(string)

	quiz, err := h.createQuizService.CreateQuiz(req.NumQuestoes, req.Dificuldade, req.Tema, userId)
	if err != nil {
		http.Error(w, "Erro ao criar quiz", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Quiz criado com sucesso",
		"quiz":    quiz,
	})
}

func (h *QuizHandler) SubmitQuiz(w http.ResponseWriter, r *http.Request) {
	var req SubmitQuizRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	resultQuiz := domain.ResultQuiz{
		QuizID:      req.QuizID,
		UserID:      req.UserID,
		Answers:     req.Answers,
		CompletedAt: time.Now(),
	}

	userId := r.Context().Value("user_id").(string)
	quiz, err := h.submitQuizService.SubmitQuiz(resultQuiz, userId)
	if err != nil {
		http.Error(w, "Erro ao enviar quiz", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Quiz enviado com sucesso",
		"quiz":    quiz,
	})
}

func (h *QuizHandler) FindManyQuizzes(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)

	quizzes, err := h.findManyQuizzesService.FindManyQuizzes(userId)
	if err != nil {
		http.Error(w, "Erro ao buscar quizzes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quizzes)
}

func (h *QuizHandler) FindQuizById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	userId := r.Context().Value("user_id").(string)

	quiz, err := h.findQuizByIdService.FindQuizById(id, userId)
	if err != nil {
		http.Error(w, "Erro ao buscar quiz", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quiz)
}
