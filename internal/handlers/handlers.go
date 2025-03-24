package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	data := []byte("Hello, World!")
	w.Write(data)
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	data := string(body)
	fmt.Printf("Received: %s\n", data)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("The data was successfully received"))
}
