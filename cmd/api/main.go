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

	dbConn, err := database.NewPostgres(cfg.DatabaseUrl)
	if err != nil {
		panic(err)
	}

	router := handler.NewRouter(dbConn)
	router.RegisterRoutes()

	fmt.Printf("Servidor rodando na porta %s\n", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, router.GetHandler()); err != nil {
		panic(err)
	}
}
