package main

import (
	"fmt"
	"goserver/internal/env"
	"goserver/internal/routes"
	"net/http"
)

func main() {
	env := env.LoadConfig()
	mux := http.NewServeMux()

	router := routes.NewRoutes()
	router.RegisterRoutes()

	fmt.Printf("Servidor rodando na porta %s\n", env.Port)

	if err := http.ListenAndServe(":" + env.Port, mux); err != nil {
		panic(err)
	}
}
