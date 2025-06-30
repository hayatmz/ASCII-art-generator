package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func WriteFile(sample string) {
	// Nom du fichier à créer
	var nomFichier string
	flag.StringVar(&nomFichier, "output", "", "result.txt")
	flag.Parse()
	// Créer un fichier
	fichier, err := os.Create(nomFichier)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return
	}
	defer fichier.Close()

	// Écrire dans le fichier
	contenu := sample
	_, err = fichier.WriteString(contenu)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}

	fmt.Printf("Le fichier %s a été créé avec succès.\n", nomFichier)
}

func main() {
	// Vérifie qu'on entre un seul argument
	arg := os.Args[1:]
	if len(arg) >= 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")

		if len(arg) < 2 {
			for _, elem := range os.Args[2] {
				if elem > 126 {
					fmt.Println("Caractères pas pris en compte.")
					return
				}
			}
		}
	}
	txt := "banners/"
	if len(os.Args)-1 >= 3 {
		input := os.Args[3]
		switch input {
		case "flame", "shadow", "thinkertoy":
			txt += input + ".txt"
		default:
			txt += "standard.txt"
		}
	}
	// Lecture du fichier
	data, err := os.ReadFile(txt)
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	// Converti data ]byte en string
	stringDeStand := string(data)

	// Converti string de standard en []string séparé par les retour à la ligne
	tabDeStand := strings.Split(stringDeStand, "\n")
	// Converti string d'argument en []string séparé par des retour à la ligne
	inputArg := strings.Split(arg[1], "\\n")

	var output string

	for _, element := range inputArg {
		if element != "" {
			for i := 1; i < 9; i++ {
				for _, element := range element {
					lettreStand := int(element) - 32
					output += tabDeStand[(lettreStand*9 + i)]
				}
				output += "\n"
			}
			output += "\n" // Ajoutez un saut de ligne après chaque élément
		} else {
			output += "\n" // Si l'élément est vide, ajoutez simplement un saut de ligne
		}
	}

	WriteFile(output) // Une fois la boucle terminée, écrivez le contenu dans le fichier
}
