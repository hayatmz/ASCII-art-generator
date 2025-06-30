package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
// Vérifie qu'on entre un seul argument
	arg := os.Args[1:]
	if len(arg) >= 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}

	for _, elem := range os.Args[1] {
		if elem > 126 {
			fmt.Println("Caractères pas pris en compte.")
			return
		}
	}

	txt := "banners/"
	if len(os.Args)-1 >= 2 {
		input := os.Args[2]
		switch input {
		case "flame" :
			txt += input + ".txt"
		case "shadow" :
			txt += input + ".txt"
		case "thinkertoy" :
			txt += input + ".txt"
		default:
			txt += "standard.txt" 
		}
	}

// Lecture du fichier
	data, err := os.ReadFile(txt)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier: %v", err)
		return
	}

// Converti data []byte en string
	stringDeStand := string(data)

// Converti string de standard en []string séparé par les retour à la ligne
	tabDeStand := strings.Split(stringDeStand, "\n")
// Converti string d'argument en []string séparé par des retour à la ligne
	inputArg := strings.Split(arg[0], "\\n")

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
			
			fmt.Print(output)
			output = ""

		} else {
			fmt.Print("\n")
		}
	}
}