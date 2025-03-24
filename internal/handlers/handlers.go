package handlers

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	data := []byte("Hello, World!")
	w.Write(data)
}
