package main

import (
	"fmt"
	"go-adv-demo/internal/hello"
	"net/http"

	"go-adv-demo/config"
)

func main() {
	conf := config.LoadConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server listen on port 8081")
	server.ListenAndServe()
}
