# ASCII-art-web
## dockerize
Ce dossier contient la version dockerisée de l'application [ASCII-art-web](https://github.com/hayatmz/ASCII-art-generator/tree/main/ascii-art-web).

C'était ma toute première expérience avec Docker.
Le Dockerfile utilisé est simple et basique, réalisé dans le but d'apprendre les concepts fondamentaux du container et de pouvoir exécuter facilement mon application Go dans un environnement isolé.

## 🐳 Description du Dockerfile

- L'image de base utilisée est l'image officielle ```golang:latest```.
- Le répertoire de travail est défini dans ```/app```.
- Tout le contenu local est copié dans ce répertoire dans le conteneur.
- L'application Go est compilée à l'intérieur du conteneur (```go build```).
- Le conteneur exécute l'application compilée au démarrage.

> Ce Dockerfile respecte les bonnes pratiques de base pour un premier Dockerfile, même s'il n'est pas optimisé.

## ⚙️ Utilisation

1. **Clone le dépôt :**<br>
```
git clone https://github.com/hayatmz/ASCII-art-generator
cd ASCII-art-generator
cd ascii-art-web
cd dockerize
```

2. **Construis l'image Docker :**<br>
```
docker build -t ascii-art-web
```

3. **Lance un conteneur Docker en exposant le port 4444 :**
```
docker run -p 4444:4444 ascii-art-web
```

5. **Accède à l'interface web :**
[http://localhost:4444](http://localhost:4444)
