package server

import (
	"net/http"

	"github.com/Kudzeri/yandex-practicum-url-shortener/internal/database"
	"github.com/Kudzeri/yandex-practicum-url-shortener/internal/handlers"
)

func NewServer() *http.Server {
	database.InitDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/", handlers.ShortURLHandler)

	port := ":8080"
	return &http.Server{
		Addr:    port,
		Handler: mux,
	}
}
