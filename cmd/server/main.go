package main

import (
	"fmt"
	"goserver/internal/db"
	"goserver/internal/platform/env"
	"goserver/internal/routes"
	"net/http"
)

func main() {
	env := env.LoadConfig()
	mux := http.NewServeMux()

	dbConn, err := db.NewPostgres(env.DatabaseUrl)
	if err != nil {
		panic(err)
	}

	router := routes.NewRoutes(dbConn)
	router.RegisterRoutes()

	fmt.Printf("Servidor rodando na porta %s\n", env.Port)

	if err := http.ListenAndServe(":"+env.Port, mux); err != nil {
		panic(err)
	}
}
