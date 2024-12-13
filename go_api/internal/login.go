package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Struct of the connexion (username + password)
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Struct of the response (TOKEN)
type LoginResponse struct {
	Token string `json:"token"`
}

// Perform the login, send the POST request, return an error in case of a problem
func PerformLogin(url, username, password string) (string, error) {
	loginRequest := LoginRequest{
		Username: username,
		Password: password,
	}

	jsonData, err := json.Marshal(loginRequest)
	if err != nil {
		return "", fmt.Errorf("erreur lors de l'encodage JSON : %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la lecture de la réponse : %v", err)
	}

	var loginResponse LoginResponse
	if err := json.Unmarshal(body, &loginResponse); err != nil {
		return "", fmt.Errorf("erreur lors du décodage JSON : %v", err)
	}

	if loginResponse.Token == "" {
		return "", fmt.Errorf("échec de la récupération du token")
	}

	return loginResponse.Token, nil
}
