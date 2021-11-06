package main

import (
	"github.com/abdullahaki/hersilia/controllers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.Hello)

	http.ListenAndServe(":8080", mux)
}
