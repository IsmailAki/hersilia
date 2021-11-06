package main

import (
	"github.com/abdullahaki/hersilia/controllers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.ReqHandler)
	mux.HandleFunc("/flush", controllers.FlushHandler)

	http.ListenAndServe(":8080", mux)
}
