package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/Kudzeri/yandex-practicum-url-shortener/internal/database"
	"github.com/Kudzeri/yandex-practicum-url-shortener/internal/utils"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	data := []byte("Hello, World!")
	w.Write(data)
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		RedirectHandler(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	url := string(body)
	shortUrl := utils.GenerateShortURL()

	_, err = database.DB.Exec("INSERT INTO urls (original_url, short_url) VALUES ($1, $2)", url, shortUrl)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	shortened := "http://localhost:8080/" + shortUrl

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortened))
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	var originalURL string
	err := database.DB.QueryRow("SELECT original_url FROM urls WHERE short_url=$1", shortURL).Scan(&originalURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}
