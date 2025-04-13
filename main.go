package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	steamID     = "76561198866021633" // Например, 76561197960287930
	steamAPIURL = "https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2/"
	appID       = 730 // CS2
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

type PlayerStats struct {
	Stats []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"stats"`
}

func main() {
	steamAPIKey, flag := os.LookupEnv("STEAM_API")
	if !flag {
		fmt.Println("STEAM_API not set")
		return
	}
	url := fmt.Sprintf("%s?key=%s&steamid=%s&appid=%d", steamAPIURL, steamAPIKey, steamID, appID)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data struct {
		Playerstats PlayerStats `json:"playerstats"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	// Вывод K/D, убийств, смертей и т. д.
	for _, stat := range data.Playerstats.Stats {
		switch stat.Name {
		case "total_kills":
			fmt.Printf("Убийства: %d\n", stat.Value)
		case "total_deaths":
			fmt.Printf("Смерти: %d\n", stat.Value)
		case "total_mvps":
			fmt.Printf("MVP: %d\n", stat.Value)
		}
	}
}
