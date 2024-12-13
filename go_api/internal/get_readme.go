package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GET the README file, could put the URL in an environment variable.
func GetReadmeContent() (string, error) {
	url := "https://raw.githubusercontent.com/SebastienGrard/go_project/testing/go_api/README.md"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la récupération du README: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erreur HTTP: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la lecture du README: %v", err)
	}

	return string(body), nil
}
