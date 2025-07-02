package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func generateAscii(str string, fileName string) string {
	// Ouvre le dossier "banners", et ajoute ".txt" pour avoir le bon fichier grâce au fileName récupérer grâce aux radios.
	data, err := os.Open("./banners/" + fileName + ".txt")
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier: %v", err)
		return ""
	}
	defer data.Close()
	// Scan le contenu du fichier et crée un tableau "catalogue" avec son contenu.
	scan := bufio.NewScanner(data)
	catalogue := make([]string, 0)
	for scan.Scan() {
		catalogue = append(catalogue, scan.Text())
	}

	tabDeStand := catalogue
	inputArg := strings.Split(str, "\\n")
	var st string
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
			st += output
			output = ""
		} else {
			st += "\n"
		}
	}
	return st
}
