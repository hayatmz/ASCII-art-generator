# ASCII-art-web
## export-file
Cette extension du projet [ascii-art-web](https://github.com/hayatmz/ASCII-art-generator/tree/main/ascii-art-web) ajoute la fonctionnalité d'exporter le résultat ASCII généré au format fichier texte téléchargeable depuis l'interface web.

## 🛠️ Fonctionnalités principales

- L'utilisateur accède à une page web avec un formulaire.
- Il entre un texte et choisit une des polices disponibles (```standard```, ```shadow```, ```thinkertoy``` ou ```flame```).
- Le texte est transformé en ASCII art et affiché sur la page.
- Génération d'un fichier texte contenant l'ASCII résultant.
- Téléchargement possible via un bouton dans l'interface.

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
cd export-file
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
