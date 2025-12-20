package routes

import (
	"goserver/internal/auth"
	"net/http"
)

type Routes struct {
	mux  *http.ServeMux
	auth *auth.AuthController
}

func NewRoutes() *Routes {
	return &Routes{
		mux:  http.NewServeMux(),
		auth: auth.NewAuthController(),
	}
}

func (r *Routes) RegisterRoutes() {
	r.auth.RegisterRoutesAuth(r.mux)
}
