package controllers

import (
	"encoding/json"
	"github.com/abdullahaki/hersilia/db"
	"net/http"
)

type Store interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type Handler struct {
	Store Store
}

func ReqHandler(w http.ResponseWriter, r *http.Request) {

	data := db.New()
	handler := &Handler{Store: data}

	switch r.Method {
	case "GET":
		handler.GetHandler(w, r)
	case "POST":
		handler.SetHandler(w, r)
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

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	value, err := h.Store.Get(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(value))
}

type setHandler struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

func (h *Handler) SetHandler(w http.ResponseWriter, r *http.Request) {
	var set setHandler
	err := json.NewDecoder(r.Body).Decode(&set)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	if set.Key == "" || set.Val == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Key and Value is required"))
		return
	}
	err = h.Store.Set(set.Key, set.Val)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.Write([]byte("OK"))
	}
}
