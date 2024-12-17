package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Server listen on port 8081")
	http.ListenAndServe(":8081", nil)
}
