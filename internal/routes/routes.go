package routes

import (
	"database/sql"
	"goserver/internal/auth"
	"net/http"
)

type Routes struct {
	mux  *http.ServeMux
	auth *auth.AuthController
}

func NewRoutes(db *sql.DB) *Routes {
	return &Routes{
		mux:  http.NewServeMux(),
		auth: auth.NewAuthController(db),
	}
}

func (r *Routes) RegisterRoutes() {
	r.auth.RegisterRoutesAuth(r.mux)
}
