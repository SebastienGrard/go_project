# Projet d’authentification en Go avec SQLite et JWT

## Description du projet
Ce projet est une application web d’authentification simple utilisant Go (Golang). Les identifiants de connexion (nom d'utilisateur et mot de passe) sont stockés dans une base de données SQLite.

## Fonctionnalités principales
- **Authentification** : Authentification des utilisateurs avec JWT (JSON Web Token).
- **Base de données** : Stockage des utilisateurs et mots de passe dans une base de données SQLite.
- **Concurrence** : Gestion des accès concurrents à la base de données à l’aide d’un mutex.

## Architecture du projet
```plaintext
├── main.go                   # Point d'entrée de l'application Go
├── internal                  # Contient les fichiers GO
│   └── handler.go            # Gestion des handlers
│   └── jwt.go                # Gestion des jetons
│   └── sql_server.go         # Gestion du serveur SQL
├── go.mod                    # Contient les paramètres GO
├── README.md                 # Documentation du projet
└── Dockerfile                # Dockerfile pour conteneuriser l’application
```

## Prérequis
- **Go** (v1.23.4 ou plus)
- **SQLite**

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

3. **Accès à l'API**
   - **Connexion** :
     ```bash
     curl -X POST http://localhost:9000/login -H "Content-Type: application/json" -d '{"username": "admin", "password": "password"}'
     ```
   - **Accès à la page de bienvenue** :
     ```bash
     curl -X GET http://localhost:9000/welcome -H "Authorization: Bearer <votre_jwt_token>"
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
docker build -t go-auth-server .
docker run -p 9000:9000 go-auth-server
```

## Améliorations futures
- **Chiffrement des mots de passe** : Remplacer le stockage direct des mots de passe par une version hachée (par exemple, en utilisant bcrypt).
- **Gestion des permissions** : Ajouter des rôles d’utilisateurs et des autorisations spécifiques.
- **Gestion des erreurs** : Meilleure gestion des erreurs HTTP et des exceptions.
- **Gestion des permissions** : Ajouter des rôles d’utilisateurs et des autorisations spécifiques.

## Auteurs
- GRARD Sébastien

## License
Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.

