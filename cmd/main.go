package main

import (
	"fmt"
	"go/adv-demo/internal/auth"
	"net/http"
)

func main() {
	//conf := config.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server listen on port 8081")
	server.ListenAndServe()
}
