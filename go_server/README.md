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

🔒 Gestion de la sécurité

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

## **Contributeurs**

GRARD Sebastien

## License

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.