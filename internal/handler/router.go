package handler

import (
	"database/sql"
	"goserver/internal/middleware"
	"net/http"
)

type Routes struct {
	mux  *http.ServeMux
	auth *AuthController
}

func NewRouter(db *sql.DB) *Routes {
	return &Routes{
		mux:  http.NewServeMux(),
		auth: NewAuthController(db),
	}
}

func (r *Routes) GetHandler() http.Handler {
	return r.mux
}

func (r *Routes) RegisterRoutes() {
	// Rotas Públicas
	r.mux.HandleFunc("POST /login", r.auth.Login)
	r.mux.HandleFunc("POST /register", r.auth.Register)

	// Rotas Protegidas
	r.mux.Handle("GET /me", middleware.AuthMiddleware(http.HandlerFunc(r.auth.Me)))
}

