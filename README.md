GO Project

MONITORING

Le projet a une finalité de monitoring.
Dans cette version, une API va :
- Se connecter sur un serveur via un login (username/password), et récupérer un token d'authentification.
- Une fois la transmission du token réalisée, elle va se connecter sur un dashboard côté client.
- Le dashboard va récupérer les informations (température CPU / GPU) de la machine, et les afficher sur le dashboard.


TODO:
Graphique
Cache ?
Stockage des données ?
SQL LITE



-----------------------------------

Initialisation de GO_API

Mise en place de la libraire :
- go get LIBRAIRIE_NAME
- go mod tidy

Utilisation de WSL + VCXSRV afin de réaliser la compilation, et l'affichage de la fenêtre.

Projet d’authentification en Go avec SQLite et JWT

🌐 Vue d’ensemble

Ce projet est une application web d’authentification simple utilisant Go (Golang). Les identifiants de connexion (nom d'utilisateur et mot de passe) sont stockés dans une base de données SQLite.

Les principales fonctionnalités sont :

Authentification des utilisateurs avec JWT (JSON Web Token).

Stockage des utilisateurs et mots de passe dans une base de données SQLite.

Gestion des accès concurrents à la base de données à l’aide d’un mutex.

📊 Architecture du projet

.
├── main.go        # Fichier principal avec les routes et la logique d'authentification
├── users.db       # Base de données SQLite (générée automatiquement au premier démarrage)
└── README.md      # Ce fichier de documentation

📘 Installation et exécution

Prérequis

Go (v1.18 ou plus)

SQLite

Installation

Clonez le dépôt :

git clone https://github.com/votre-utilisateur/votre-depot.git
cd votre-depot

Exécutez l'application :

go run main.go

Accédez à l'application sur :
http://localhost:8080

🔧 Utilisation de l’API

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

🔠 Détails techniques

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

🔄 Améliorations possibles

Chiffrement des mots de passe : Remplacer le stockage direct des mots de passe par une version hachée (par exemple, en utilisant bcrypt).

Gestion des permissions : Ajouter des rôles d’utilisateurs et des autorisations spécifiques.

Gestion des erreurs : Meilleure gestion des erreurs HTTP et des exceptions.

📊 Diagrammes et illustrations

🔄 Flux de l'authentification

graph TD;
    A[Démarrage de l'application] --> B[Requête POST /login];
    B -->|Requête valide| C{Vérification de l'utilisateur};
    C -->|Valide| D[Génération du JWT];
    D --> E[Renvoi du token JWT];
    C -->|Invalide| F[Erreur 401: Non autorisé];

🌐 Structure des routes

GET  /welcome     - Accès restreint, nécessite un token JWT
POST /login       - Génère un JWT en cas de connexion réussie

🌐 Contributeurs

GRARD Sebastien

N’hésitez pas à proposer des suggestions ou des améliorations via des issues ou des pull requests.

