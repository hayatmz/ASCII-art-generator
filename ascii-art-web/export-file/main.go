package main

import (
	"fmt"
	"net/http"
)

const port = ":4447"

func main() {
	// Crée un handler et y assigne une fonction.
	http.HandleFunc("/", ErrorHandler)
	http.HandleFunc("/ascii-art", AsciiHandler)
	http.HandleFunc("/download", downloadHandler)

	// ListenAndServe écoute le code et sert sur le serveur.
	fmt.Println("(http://localhost:4447) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
