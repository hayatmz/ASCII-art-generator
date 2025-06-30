package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Fonction qui récupère la taille du terminal.
func terminalSize() int {
// exec.Command appelle la commande "stty size"
	cmd := exec.Command("stty", "size")
// cmd.Stdin pour définir l'entrée de la commande sur le contenu d'un fichier
	cmd.Stdin = os.Stdin
// cmd.Output recupère le résultat de la commande
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
// Transformation du résultat de la commande []byte en []string
	s := string(out)
	s = strings.TrimSpace(s)
	strtab := strings.Split(s, " ")
	
// Renvoie la valeur de la longueur(string) en int. ([0] = hauteur & [1] = longueur)
	width, err := strconv.Atoi(strtab[1])
	if err != nil {
		log.Fatal(err)
	}
	return width
}

func main() {
	arg := os.Args[1:]
	txt := "banners/standard.txt"

// Vérifie qu'on entre pas plus de 3 arguments.
	if len(arg) >= 4 || len(arg) <= 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=right something standard")
		return
	}

// Vérifie que l'arg[0] --align=<type> soit bien écrit.
	if arg != nil {
		if arg[0] != "--align=right" && arg[0] != "--align=left" && arg[0] != "--align=center" && arg[0] != "--align=justify" {
			fmt.Println("Usage [OPTION]: --align=right, --align=left, --align=center or --align=justify.")
			return
		}

		if len(arg) == 3 {
			if arg[2] != "standard" && arg[2] != "shadow" && arg[2] != "thinkertoy" && arg[2] != "flame" {
			fmt.Println("Usage [BANNER]: standard, shadow, thinkertoy or flame.")
			return
			}

		// Récupère le bon fichier du template à utiliser.
			switch os.Args[3] {
			case "shadow" :
				txt = "banners/shadow.txt"
			case "standard" :
				txt = "banners/standard.txt"
			case "thinkertoy" :
				txt = "banners/thinkertoy.txt"
			case "flame" :
				txt = "banners/flame.txt"
			}
		}
	}

// S'assure qu'on entre uniquement des caractères compris dans la table ascii.
	for _, elem := range os.Args[2] {
		if elem < 32 || elem > 126 {
			fmt.Println("Caractères pas pris en compte.")
			return
		}
	}

// Lecture du fichier du template.
	data, err := os.ReadFile(txt)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier: %v\n", err)
		return
	}

// Converti contenu du fichier en strin, puir contenu du fichier et d'arg en []string, séparé par des retour à la ligne.
	stringDeTxt := string(data)
	tabTxt := strings.Split(stringDeTxt, "\n")

	inputArg := strings.Split(arg[1], "\\n")
	var output string
	width := terminalSize()
	var lettresAscii int

// Range sur le string passé en argument pour le transformer en représentation ASCII à partir du fichier txt.
	for _, elem := range inputArg {

		if elem == "" {
			fmt.Println("Argument vide.\nUsage: go run . [OPTION] [STRING] [BANNER]\nOu: go run . [OPTION] [STRING]")
		} else {
			for i := 1; i < 9; i++ {
				var pad bool = false
				var motJusti string
				for _, element := range elem {
					if os.Args[1] != "--align=justify" || element != ' ' { 
						lettresAscii += len(tabTxt[(int(element)-32)*9 + i])
					}
					output += tabTxt[(int(element)-32)*9 + i]
				}
				espacesCompteur := len(strings.Fields(elem)) - 1

				if len(output) > width {
					fmt.Printf("Taille de l'argument supérieure à celle du terminal(%v).\n", width)
					return
				}

				tabarg := strings.Split(os.Args[2], " ")
				if os.Args[1] == "--align=justify" && len(tabarg) == 1 {
					fmt.Println("Taille de l'argument incompatible avec --align=justify.\nAjoute un mot ou essaye avec --align=right, --align=left ou --align=center.")
					return
				}

				if os.Args[1] == "--align=justify" {
					for _, element := range elem {
						if element == ' ' {
							fmt.Print(motJusti)
							if ((width-lettresAscii) % espacesCompteur) != 0 && !pad {
								fmt.Print(" ")
								pad = true
							}
							fmt.Print(strings.Repeat(" ", (width-lettresAscii) / espacesCompteur))
							motJusti = ""
							
						} else {
							motJusti += tabTxt[(int(element)-32)*9 + i]
						}
					}
				}

			espaces := (int(float64(width-lettresAscii)))
			switch os.Args[1] {
			case "--align=right" :
				droite := strings.Repeat(" ", espaces) + output
				fmt.Println(droite)
			case "--align=left" :
				gauche := (output + strings.Repeat(" ", espaces))
				fmt.Println(gauche)
			case "--align=center" :
				centre := strings.Repeat(" ", espaces/2) + output + strings.Repeat(" ", espaces/2)
				fmt.Println(centre)
			case "--align=justify" :
				fmt.Print(motJusti)
			}
			output = ""
			lettresAscii = 0
			}
		}
	}
}