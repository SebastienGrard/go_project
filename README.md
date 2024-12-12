# Projet d’authentification en Go avec SQLite et JWT

## Description du projet
Ce projet est une application web d’authentification simple utilisant Go (Golang). Les identifiants de connexion (nom d'utilisateur et mot de passe) sont stockés dans une base de données SQLite.

## Fonctionnalités principales
- **Authentification** : Authentification des utilisateurs avec JWT (JSON Web Token).
- **BDD** : Stockage des utilisateurs et mots de passe dans une base de données SQLite.
- **Concurrence** : Gestion des accès concurrents à la base de données à l’aide d’un mutex.

## Architecture du projet
```plaintext
├── README.md                 # Documentation du projet
├── internal                  # Contient les fichiers GO
│   └── handler.go            # Gestion des handlers
│   └── jwt.go                # Gestion des jetons
│   └── sql_server.go         # Gestion du serveur SQL
├── main.go                   # Point d'entrée de l'application Go
├── go.mod                    # Contient les paramètres GO
└── Dockerfile                # Dockerfile pour conteneuriser l’application
```

## Prérequis
- **Go** (v1.23.4 ou plus)
- **SQLite** 

## Installation :
## Installation et exécution
1. **Clôner le dépôt**
   ```bash
   git clone https://github.com/SebastienGrard/go_project.git
   cd go_projet/go_server
   ```
2. **Démarrer l'application**
   ```bash
   go run main.go
   ```
   L'application sera disponible à l'adresse suivante : [http://localhost:9000](http://localhost:9000)

## **Utilisation de l’API**

1. Route : /login (POST)

Description : Authentification de l'utilisateur et obtention d'un JWT.

Requête :

{
  "username": "votre_nom_utilisateur",
  "password": "votre_mot_de_passe"
}

Réponse :

{
  "token": "votre_jwt_token"
}

2. Route : /welcome (GET)

Description : Accès à la page de bienvenue (authentification nécessaire).

Requête :

GET /welcome
Authorization: Bearer <votre_jwt_token>

Réponse :

{
  "message": "Bienvenue utilisateur"
}

## **Détails techniques**

1. Base de données SQLite

Le fichier users.db est créé automatiquement au premier démarrage.

La table users est créée si elle n’existe pas.

Structure de la table :

CREATE TABLE users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  username TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
);

2. Gestion des routes

Route

Méthode

Description

/login

POST

Authentification et obtention d’un JWT

/welcome

GET

Accès à la page de bienvenue

3. Gestion des JWT

Un token JWT est généré lors de la connexion.

Le token inclut les claims suivants :

{
  "sub": "username",
  "exp": 1699999999 // Date d'expiration
}

La clé secrète de signature du JWT est :

var jwtKey = []byte("secret_key")

## **Gestion de la sécurité**

Chiffrement des mots de passe :

Amélioration potentielle : utilisation de bcrypt au lieu de stocker les mots de passe en clair.

Vérification du JWT :

Le JWT est vérifié à chaque requête à /welcome.

## **Améliorations possibles**

- **Chiffrement des mots de passe** : Remplacer le stockage direct des mots de passe par une version hachée (par exemple, en utilisant bcrypt).
- **Gestion des permissions** : Ajouter des rôles d’utilisateurs et des autorisations spécifiques.
- **Gestion des erreurs** : Meilleure gestion des erreurs HTTP et des exceptions.

## **Flux de l'authentification**

graph TD;
    A[Démarrage de l'application] --> B[Requête POST /login];
    B -->|Requête valide| C{Vérification de l'utilisateur};
    C -->|Valide| D[Génération du JWT];
    D --> E[Renvoi du token JWT];
    C -->|Invalide| F[Erreur 401: Non autorisé];

## **Structure des routes**

GET  /welcome     - Accès restreint, nécessite un token JWT
POST /login       - Génère un JWT en cas de connexion réussie


## **TEST**

curl -X POST http://localhost:9000/login -H "Content-Type: application/json" -d '{"username": "admin", "password": "password"}'

curl -X GET http://localhost:9000/welcome -H "Authorization: Bearer YOUR_TOKEN_ID"

##############################################################################################
##############################################################################################
##############################################################################################
##############################################################################################
##############################################################################################

# Projet avec GOLANG

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
