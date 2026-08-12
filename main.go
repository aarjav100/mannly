package main

import (
	"net/http"
	"url-shortener/handlers"
	"url-shortener/storage"
)

func main() {
	store := storage.NewStore()

	handler := handlers.NewHandler(store)

	http.HandleFunc("/shorten", handler.ShortenURL)
	http.HandleFunc("/", handler.RedirectURL)

	port := ":8080"
	println("Server is running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		println("Error starting server:", err)
	}

}
