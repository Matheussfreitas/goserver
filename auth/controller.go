package auth

import (
	"encoding/json"
	"errors"
	"net/http"
)

type AuthController struct {
	authService *AuthService
}

// NewAuthController atua como o 'construtor'
func NewAuthController() *AuthController {
	return &AuthController{authService: NewAuthService()}
}

// Exemplo de método da 'classe'
// RegisterRoutes registra todas as rotas de autenticação
func (c *AuthController) RegisterRoutes(mux *http.ServeMux) {
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

	if err := c.authService.Login(req.Email, req.Password); err != nil {
		if errors.Is(err, ErrUserNotFound) {
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
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login realizado com sucesso",
		"token":   "exemplo-de-token-jwt",
	})
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	if err := c.authService.Register(req.Email, req.Password); err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
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
