package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

// Lecture du fichier de bannière contenant l'art ASCII (standard.txt)
	data, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier: %v", err)
		return
	}

// Convertit le contenu du fichier ([]byte) en string
	stringDeStand := string(data)

// Découpe le contenu en lignes (chaque caractère ASCII du fichier a une hauter de 8 lignes + 1 ligne vide entre les blocs)
	tabDeStand := strings.Split(stringDeStand, "\n")

// Récupère l'argument passé en commande et découpe à chaque "\n" pour gérer les sauts de ligne logiques
	inputArg := strings.Split(os.Args[1], "\\n")

var output string

	// Parcourt chaque ligne à afficher
	for _, ligne := range inputArg {
		if ligne != "" {
			// Pour chaque ligne, on affiche l'art ASCII (de 1 à 8, pour correspondre au nombre de lignes de chaque caractère)
			for i := 1; i < 9; i++ {
				for _, ligne := range ligne {
					// Calcul de l'index dans le tableau ASCII
					lettreStand := int(ligne) - 32 // Les caractères ASCII commencent à l'index 32
					output += tabDeStand[(lettreStand*9 + i)]
				}
				output += "\n"
			}
			
			fmt.Print(output)
			output = ""

		} else {
			// Si la ligne est vide (double \n), on imprime une ligne vide
			fmt.Print("\n")
		}
	}
}