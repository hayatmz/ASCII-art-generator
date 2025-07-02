package main

import (
	"fmt"
	"net/http"
)

const port = ":4444"

func main() {
	// Crée un handler et y assigne une fonction.
	http.HandleFunc("/", ErrorHandler)
	http.HandleFunc("/ascii-art", AsciiHandler)

	// ListenAndServe écoute le code et sert sur le serveur.
	fmt.Println("(http://localhost:4444) - Server started on port", port)
	http.ListenAndServe(port, nil)
}