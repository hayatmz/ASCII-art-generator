# ASCII-art-generator
ASCII-art-generator est une collection de programmes écrits en Go qui transforment du texte standard en représentation graphique stylisée utilisant des caractères ASCII.<br>
Le projet s'articule autour de différentes variantes, allant d'une version simple en ligne de commande à des extensions plus avancées (choix de police, aligenement, export dans un fichier...).
>Deuxième projet crée en Décembre 2023 au sein de la formation de [Zone01](https://campus-saint-marc.com/zone-01/).

## 🎯 Objectifs pédagogiques
- Comprendre la manipulation de fichiers en Go.
- Travailler avec des arguments CLI.
- Travailler avec les chaînes de caractères.
- Etendre une base de code vers des variantes avancées.

## 🔧 Fonctionnalités principales

- **[ascii-art/main.go](./ascii-art/main.go)**<br>
Version de base : prend une chaîne en argument et l'affiche en ASCII dans le terminal.<br>
_Utilise la bannière ```standard.txt``` située dans le dossier ```banners/```_<br>

## 📦 Extensions / Sous-dossiers
- **[fs](./ascii-art/fs/main.go)**
Ajoute la possibilité de choisir la bannière utilisée pour la conversion ASCII.

    📚 Bannières disponibles dans le dossier ```banners/```:<br>
    - ```standard.txt```<br>
    - ```shadow.txt```<br>
    - ```thinkertoy.txt```<br>
    - ```flame.txt``` (_créée manuellement_)<br>

- **[justify](./ascii-art/justify/main.go)**
A le même fonctionnement que [fs](./ascii-art/fs/main.go), avec un troisième argument pour gérer l'alignement du texte : ```left```, ```right```, ```center```, ou ```justify```.

- **[output](./ascii-art/output/main.go)**
A également le même fonctionnement que [fs](./ascii-art/fs/main.go), avec une extension permettant d'enregistrer le résultat dans un fichier texte.

## 🚀 Utilisation

1. **Assure toi d'avoir Go installé** sur ta machine :<br>
```
go version
```
Si ce n'est pas le cas, [installe Golang](https://go.dev/doc/install).

2. **Clone le dépôt :**<br>
```
git clone https://github.com/hayatmz/ASCII-art-generator
cd ASCII-art-generator
```

3. **Lance le programme :**

- Version de base :
```
go run main.go "texte de ton choix"
```
**ou** avec un seul mot :<br>
```
go run main.go motDeTonChoix
```
> Pour chaque projet, la règle du mot (sans les "&nbsp;") ou "texte" (avec les "&nbsp;") sera appliquée.

- fs (choix de bannière):<br>
```
go run fs/main.go motDeTonChoix flame
```
> ```flame```, ```shadow```, ```standard```, ou ```thinkertoy```.

- justify (alignement) :<br>
```
go run justify/main.go --align=left Mot nomDeBanner
```
> ```--align=left```, ```--align=right```, ```--align=center```, ou ```--align=justify```.

- Output (création de fichier) :
```
go run output/main.go --output="NomFichier" "texte de ton choix" nomDeBanner
```