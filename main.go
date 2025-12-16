package main

import (
	"fmt"
	"goserver/auth"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	controller := auth.NewAuthController()
	controller.RegisterRoutes(mux)

	fmt.Println("Servidor rodando na porta 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
