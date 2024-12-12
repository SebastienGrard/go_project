# Projet de Monitoring de VM avec Go

## Description du projet
Ce projet consiste à développer une API en Go permettant de surveiller les performances d’une machine virtuelle (VM) hébergée sur Azure. Le système inclut un portail web avec un tableau de bord interactif, un système de connexion sécurisé par JWT, et une gestion de cache pour optimiser les performances de l'application.

## Fonctionnalités principales
- **Connexion sécurisée** : Authentification via JWT.
- **Tableau de bord dynamique** : Visualisation des performances en temps réel (CPU, RAM, disque, GPU).
- **Cache de réponse** : Le contenu de la page de bienvenue est mis en cache pendant 2 minutes pour optimiser les performances.
- **Notifications d'alerte** : Alertes par e-mail en cas de dépassement des seuils prédéfinis.

## Architecture du projet
```plaintext
├── main.go                   # Point d'entrée de l'application Go
├── templates                 # Contient les fichiers HTML
│   └── welcome.html          # Page de bienvenue avec graphiques dynamiques
├── README.md                 # Documentation du projet
├── .env                      # Variables d'environnement (non inclus dans le dépôt)
└── Dockerfile                # Dockerfile pour conteneuriser l’application
```

## Prérequis
- **Go** (v1.20+)
- **Docker** (pour exécuter l’application dans un conteneur)
- **Navigateur web** (pour accéder au tableau de bord)

## Installation et exécution
1. **Clôner le dépôt**
   ```bash
   git clone https://github.com/votre-utilisateur/votre-projet.git
   cd votre-projet
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
                                 │
                                 └-> [Serve le contenu depuis le cache]
```

## Aperçu des graphiques (Tableau de bord)

Le tableau de bord contient 4 graphiques interactifs :
- **Utilisation du CPU**
- **Utilisation de la RAM**
- **Utilisation du disque**
- **Température du GPU**

![Aperçu des graphiques](https://via.placeholder.com/800x400.png?text=Graphiques+du+tableau+de+bord)

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

## Sécurité
- Les mots de passe doivent être chiffrés.
- Les clés JWT doivent être protégées via des variables d’environnement.

## Améliorations futures
- Ajout de la persistance des données dans une base de données (PostgreSQL).
- Mise en place d’un système de rôles d’utilisateurs (admin, utilisateur).
- Mise en place d’alertes e-mail en cas de dépassement des seuils.

## Auteurs
- GRARD Sébastien - Développeur principal

## License
Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.


IDEE Amelioration :
Rentrer les login en variable d'environnement (via docker run)
