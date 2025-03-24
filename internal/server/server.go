package server

import (
	"net/http"

	"github.com/Kudzeri/yandex-practicum-url-shortener/internal/handlers"
)

func NewServe() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handlers.HelloHandler)

	port := ":8080"
	return &http.Server{
		Addr:    port,
		Handler: mux,
	}
}
