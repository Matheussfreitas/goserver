package handler

import (
	"database/sql"
	"encoding/json"
	"goserver/internal/repository"
	"goserver/internal/service/quizzes"
	"net/http"
)

type QuizHandler struct {
	findManyQuizzesService *quizzes.FindManyQuizzesService
	createQuizService      *quizzes.CreateQuizService
}

func NewQuizHandler(db *sql.DB) *QuizHandler {
	repo := repository.NewQuizRepository(db)
	return &QuizHandler{
		findManyQuizzesService: quizzes.NewFindManyQuizzesService(repo, db),
		createQuizService:      quizzes.NewCreateQuizService(repo, db),
	}
}

type CreateQuizRequest struct {
	Tema        string `json:"tema"`
	NumQuestoes int    `json:"numQuestoes"`
	Dificuldade string `json:"dificuldade"`
}

func (h *QuizHandler) CreateQuiz(w http.ResponseWriter, r *http.Request) {
	var req CreateQuizRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	quiz, err := h.createQuizService.CreateQuiz(req.Tema, req.NumQuestoes, req.Dificuldade)
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
