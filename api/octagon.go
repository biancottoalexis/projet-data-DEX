package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"projet-data-DEX/models"
)

func GetFighter(id string) (models.Fighter, error) {
	var fighter models.Fighter

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	url := "https://api.octagon-api.com/fighter/" + id

	resp, err := client.Get(url)

	if err != nil {
		return fighter, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fighter, fmt.Errorf("erreur API : %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&fighter)

	return fighter, err
}

func GetAllFighters() (map[string]models.Fighter, error) {
	var fighters map[string]models.Fighter

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://api.octagon-api.com/fighters")

	if err != nil {
		return fighters, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fighters, fmt.Errorf("erreur API : %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&fighters)

	return fighters, err
}


