package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 1️⃣ Connexion au serveur et récupération du token
	loginURL := "http://localhost:8080/login"
	loginPayload := map[string]string{
		"username": "admin",
		"password": "password",
	}

	payloadBytes, _ := json.Marshal(loginPayload)
	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Erreur de connexion:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var responseMap map[string]string
	json.Unmarshal(body, &responseMap)

	if token, exists := responseMap["token"]; exists {
		fmt.Println("✅ Token obtenu :", token)

		// 2️⃣ Accéder au dashboard avec le token
		welcomeURL := "http://localhost:8080/welcome"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", welcomeURL, nil)
		req.Header.Add("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Erreur lors de l'accès au dashboard:", err)
			return
		}
		defer resp.Body.Close()

		welcomeBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("📢 Réponse du serveur:", string(welcomeBody))
	} else {
		fmt.Println("❌ Échec de la connexion :", string(body))
	}
}
