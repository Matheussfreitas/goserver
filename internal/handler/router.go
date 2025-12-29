package handler

import (
	"database/sql"
	"net/http"
)

type Routes struct {
	mux  *http.ServeMux
	auth *AuthController
}

func PublicRoutes(db *sql.DB) *Routes {
	return &Routes{
		mux:  http.NewServeMux(),
		auth: NewAuthController(db),
	}
}

func (r *Routes) AuthRoutes() {
	r.auth.RegisterRoutesAuth(r.mux)
}
