package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"goserver/internal/repository"
	"goserver/internal/service"
	"net/http"
)

type AuthController struct {
	authService *service.AuthService
}

// NewAuthController atua como o 'construtor'
func NewAuthController(db *sql.DB) *AuthController {
	repo := repository.NewUserRepository(db)
	return &AuthController{authService: service.NewAuthService(db, repo)}
}

// Exemplo de método da 'classe'
// RegisterRoutesAuth registra todas as rotas de autenticação
func (c *AuthController) RegisterRoutesAuth(mux *http.ServeMux) {
	mux.HandleFunc("POST /login", c.Login)
	mux.HandleFunc("POST /register", c.Register)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	user, token, err := c.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Erro ao fazer login",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login realizado com sucesso",
		"user":    user,
		"token":   token,
	})
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	if err := c.authService.Register(r.Context(), req.Email, req.Password); err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			w.WriteHeader(http.StatusConflict) // 409 Conflict
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Erro ao fazer cadastro",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Usuário criado com sucesso",
	})
}
