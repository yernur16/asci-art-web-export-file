package main

import (
	"log"
	"net/http"

	"ascii/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/ascii-art", handlers.Ascii)
	mux.HandleFunc("/about.html", handlers.About)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))
	log.Println("Starting the server http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

