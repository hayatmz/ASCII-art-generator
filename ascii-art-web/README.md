# ASCII-art-web
ASCII-art-web est une extension web du projet [ascii-art](../ascii-art/), dévelopée en Go en utilisant uniquement les packages standards. Ce projet permet de convertir du texte en art ASCII via une interface web simple, sans JavaScript.

Ce projet principal est enrichi par deux extensions situées dans des sous-dossiers :
- [export-file](./export-file/)
- [dockerize](./dockerize/)

> Projet crée début 2024, dans le cadre de la formation [Zone01](https://campus-saint-marc.com/zone-01/). Ce fut **ma toute première expérience avec le développement web**. L'interface a donc été réalisée avec mes compétences débutantes.

## 💡 Objectifs pédagogiques
- Comprendre et utiliser le protocole HTTP (```GET```, ```POST```, gestion des erreurs).
- Créer une interface web simple avec ```HTML``` et ```CSS```.
- Gérer des fichiers statiques.
- Apprendre l'utilisation des en-têtes HTTP pour le téléchargement de fichiers.
- Dockeriser une application Go web.

## 🛠️ Fonctionnalités principales

- L'utilisateur accède à une page web avec un formulaire.
- Il entre un texte et choisit une des polices disponibles (```standard```, ```shadow```, ```thinkertoy``` ou ```flame```).
- Le texte est transformé en ASCII art et affiché sur la page.
- Des pages personnalisées sont affichées en cas d'erreurs :
    - ```404``` : Not Found
    - ```400``` : Bad Request (texte invalide, police inconnue, etc.)
    - ```405``` : Method Not Allowed

## 📦 Extensions / Sous-dossiers
- **[export-file](./export-file/)**
Ajoute la possibilité d'exporter le résultat ASCII généré au format fichier texte via un bouton sur l'interface web.

- **[dockerize](./dockerize/)**
Dockerise l'application web afin de l'exécuter dans un conteneur léger, transportable et propre.

> Utilisation dans les README de chaque dossier.

## ⚙️ Utilisation

1. **Assure toi d'avoir Go installé** sur ta machine :<br>
```
go version
```
Si ce n'est pas le cas, [installe Golang](https://go.dev/doc/install).

2. **Clone le dépôt :**<br>
```
git clone https://github.com/hayatmz/ASCII-art-generator
cd ASCII-art-generator
cd ascii-art-web
```

3. **Initialise le module Go :**<br>
```
go mod init ascii-art-web
go mod tidy
```

4. **Lance le serveur :**
```
go run .
```

5. **Accède à l'interface web :**
[http://localhost:4444](http://localhost:4444)
