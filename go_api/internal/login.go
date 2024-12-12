package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Structure de la requête de connexion
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Structure de la réponse avec le token
type LoginResponse struct {
	Token string `json:"token"`
}

func PerformLogin(url, username, password string) (string, error) {
	loginRequest := LoginRequest{
		Username: username,
		Password: password,
	}

	// Encodage JSON de la requête
	jsonData, err := json.Marshal(loginRequest)
	if err != nil {
		return "", fmt.Errorf("erreur lors de l'encodage JSON : %v", err)
	}

	// Envoi de la requête POST
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	// Lecture de la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la lecture de la réponse : %v", err)
	}

	// Décodage du JSON dans LoginResponse
	var loginResponse LoginResponse
	if err := json.Unmarshal(body, &loginResponse); err != nil {
		return "", fmt.Errorf("erreur lors du décodage JSON : %v", err)
	}

	// Vérification du token
	if loginResponse.Token == "" {
		return "", fmt.Errorf("échec de la récupération du token")
	}

	return loginResponse.Token, nil
}
