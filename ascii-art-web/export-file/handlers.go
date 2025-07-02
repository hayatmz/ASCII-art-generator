package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
)

type asciiArt struct {
	textarea string
}

var ascii asciiArt

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	fileSystem := http.Dir("./templates")
	fileServer := http.FileServer(fileSystem)
	_, err := fileSystem.Open(path.Clean(r.URL.Path))

	// S'il ne trouve pas "./templates", que l'url est différent de / ou /ascii-art, je renvoie sur la page html 404.
	if os.IsNotExist(err) && r.URL.Path != "/ascii-art" && r.URL.Path != "/" {
		http.Redirect(w, r, "/404Page.html", http.StatusSeeOther)
		return
	}

	fileServer.ServeHTTP(w, r)
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {

	// Si le lien ne renvoie pas sur la page html contenant le formulaire "post", je renvoie sur la page d'erreur methodPage.
	if r.Method != "POST" {
		http.Redirect(w, r, "/methodPage.html", http.StatusSeeOther)
		return
	}
	r.ParseForm()

	// Si le traitement de l'input de mon textarea contient un caractère non supporté par l'ascii-art, je renvoie sur la page d'erreur 400.
	InputTexte := r.FormValue("inputTxt")
	for _, elem := range InputTexte {
		if !(elem >= 32 && elem <= 126) && elem != 10 {
			http.Redirect(w, r, "/400Page.html", http.StatusSeeOther)
			return
		}
	}
	// Je récupère le fichier sélectionné par les radios au nom "banner".
	banner := r.FormValue("banner")
	tmpl, err := template.ParseFiles("./templates/asciiPage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Je renvoie mes arguments inputTexte et banner récupérés dans mon formulaire html vers mon programme qui génère l'ascii art.
	ascii.textarea = generateAscii(InputTexte, banner)
	tmpl.Execute(w, ascii.textarea)

	// Enregistrement du résultat dans un fichier texte
	err = ioutil.WriteFile("./templates/resultat_ascii.txt", []byte(ascii.textarea), 0644)
	if err != nil {
		// Handle the error appropriately
		fmt.Println("Erreur lors de l'écriture du fichier :", err)
		return
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {

	content := []byte(ascii.textarea)
	fmt.Println(r.Method)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename=\"ascii.txt\"")
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))

	w.WriteHeader(http.StatusOK)
	_, err := w.Write(content)
	if err != nil {
		fmt.Println("Error writing response:", err)
	}

}