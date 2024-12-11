package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project-root/internal/auth"
	"project-root/internal/services"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	json.NewDecoder(r.Body).Decode(&loginRequest)

	// Simuler la validation de l'utilisateur (à remplacer par une BDD)
	if loginRequest.Username == "admin" && loginRequest.Password == "password" {
		jwtToken, _ := auth.GenerateJWT(loginRequest.Username)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Token: %s", jwtToken)))
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	metrics, err := services.CollectMetrics()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(metrics)
}
