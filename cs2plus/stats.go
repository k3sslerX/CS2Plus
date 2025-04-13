package cs2plus

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

const (
	steamAPIURL = "https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2/"
	appID       = 730 // CS2
)

type PlayerStats struct {
	Stats []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"stats"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func getApiKey() (string, error) {
	steamAPIKey, flag := os.LookupEnv("STEAM_API")
	if !flag {
		fmt.Println("STEAM_API not set")
		return "", errors.New("STEAM_API not set")
	}
	return steamAPIKey, nil
}

func GetOverallStats(user *SteamUser) {
	steamAPIKey, err := getApiKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	url := fmt.Sprintf("%s?key=%s&steamid=%d&appid=%d", steamAPIURL, steamAPIKey, user.SteamID64, appID)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var data struct {
		PlayerStats PlayerStats `json:"playerstats"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
}
