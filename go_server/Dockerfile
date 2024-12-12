# 1️⃣ Étape de construction
FROM golang:1.23.4 AS builder

# Activer CGO pour SQLite et spécifier l'architecture de compilation
ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# Installer les dépendances nécessaires pour CGO (GCC, etc.)
RUN apt-get update && apt-get install -y \
    build-essential \
    gcc \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier le code source Go dans le conteneur
COPY . .

# Télécharger les dépendances du projet
RUN go mod tidy

# Compiler l'application
RUN go build -o main .

# 2️⃣ Étape d'exécution
FROM debian:12

# Installer les bibliothèques nécessaires pour exécuter l'application
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier l'application compilée depuis l'étape de construction
COPY --from=builder /app/main .

# Copier la base de données SQLite (optionnel)
COPY --from=builder /app/users.db .

# Exposer le port sur lequel l'application écoute
EXPOSE 9000

# Commande pour exécuter l'application
CMD ["./main"]
