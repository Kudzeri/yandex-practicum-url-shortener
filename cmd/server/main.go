package main

import (
	"log"

	"github.com/Kudzeri/yandex-practicum-url-shortener/internal/server"
)

func main() {
	srv := server.NewServer()

	log.Println("Server started on http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
