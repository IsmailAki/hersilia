package controllers

import "net/http"

func ReqHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetHandler(w, r)
	case "POST":
		SetHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}

func FlushHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("Flushing"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET"))
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST"))
}
