package main

import (
	"fmt"
	"goserver/internal/config"
	"goserver/internal/database"
	"goserver/internal/handler"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	mux := http.NewServeMux()

	dbConn, err := database.NewPostgres(cfg.DatabaseUrl)
	if err != nil {
		panic(err)
	}

	router := handler.PublicRoutes(dbConn)
	router.AuthRoutes()

	fmt.Printf("Servidor rodando na porta %s\n", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		panic(err)
	}
}
