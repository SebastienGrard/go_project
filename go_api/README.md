# GO API

## Description du projet
Ce projet consiste à développer un serveur d'authentification. Il utilise une base de données SQLite. L'utilisateur se connecte via ses identifiants, et récupère un jeton pour accéder à des services.

## Fonctionnalités principales
- **Connexion sécurisée** : Authentification via JWT.
- **Affichage du README.md** : Visualisation du REAMDE.md du projet, avec le texte formaté.

## Architecture du projet
```plaintext
├── main.go                   # Point d'entrée de l'application Go
├── templates                 # Contient les fichiers HTML
│   └── welcome.html          # Page de bienvenue avec graphiques dynamiques
├── README.md                 # Documentation du projet
├── internal                  # Contient les fichiers GO
│   └── browser.go            # Gestion de l'ouverture du browser
│   └── converto_to_html.go   # Conversion du README et mise en page
│   └── get_readme.go         # Recuperation du README
│   └── login.go              # Gestion du Token et du login
│   └── server.go             # Gestion du serveur
└── Dockerfile                # Dockerfile pour conteneuriser l’application
```

## Prérequis
- **Go** (v1.23.4)
- **Docker** (pour exécuter l’application dans un conteneur)
- **Navigateur web** (pour accéder à la page web)

## Installation et exécution
1. **Clôner le dépôt**
   ```bash
   git clone https://github.com/SebastienGrard/go_project.git
   cd go_projet/go_api
   ```
2. **Démarrer l'application**
   ```bash
   go run main.go
   ```
   L'application sera disponible à l'adresse suivante : [http://localhost:8080](http://localhost:8080)

3. **Accès à l'API**
   - **Connexion** :
     ```bash
     curl -X POST http://localhost:8080/login -d '{"username":"admin", "password":"password"}'
     ```
   - **Accès à la page de bienvenue** :
     ```bash
     curl -H "Authorization: Bearer <TOKEN>" http://localhost:8080/welcome
     ```

## Exemple de flux de connexion
1. L'utilisateur se connecte avec un nom d'utilisateur et un mot de passe.
2. L'API retourne un JWT.
3. L'utilisateur utilise le JWT pour accéder à la page de bienvenue.

## Infographie du projet

**1. Authentification JWT**
```
[Utilisateur] -- (username/password) --> [API] -- (vérifie) --> [Génère JWT] -- (retourne JWT) --> [Utilisateur]
```

**2. Accès au contenu avec cache**
```
[Utilisateur] -- (demande) --> [API] -- (vérifie JWT) --> [Cache]
                              
```


## Variables d'environnement
Créez un fichier `.env` à la racine du projet avec les clés suivantes :
```
JWT_SECRET=super_secret_key
```

## Docker
Pour exécuter l’application dans un conteneur Docker :
```bash
docker build -t go-vm-monitoring .
docker run -p 8080:8080 go-vm-monitoring
```

## Améliorations futures
- Ajout de la persistance des données dans une base de données (PostgreSQL).
- Mise en place d’un système de rôles d’utilisateurs (admin, utilisateurs)
- Les mots de passe doivent être chiffrés.
- Les mots de passe peuvent être initialisés via un fichier.
- Les clés JWT doivent être protégées via des variables d’environnement.

## Auteurs
- GRARD Sébastien

## License
Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.
