package statsValve

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	overallAPI = "https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2/"
	appID      = 730 // CS2
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

func GetOverallStats(user *SteamUser) (*PlayerStats, error) {
	steamAPIKey, err := getApiKey()
	if err != nil {
		return nil, err
	}
	steamUrl := fmt.Sprintf("%s?key=%s&steamid=%s&appid=%d", overallAPI, steamAPIKey, user.SteamID64, appID)
	resp, err := http.Get(steamUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data struct {
		PlayerStats PlayerStats `json:"playerstats"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data.PlayerStats, nil
}

func GetOverallAccuracy(user *SteamUser) (map[string]float64, error) {
	playerStats, err := GetOverallStats(user)
	if err == nil {
		shoots := make(map[string]int, 53)
		hits := make(map[string]int, 53)
		accuracies := make(map[string]float64, 53)
		for _, stat := range playerStats.Stats {
			if strings.HasPrefix(stat.Name, "total_shots_") {
				weapon := strings.TrimPrefix(stat.Name, "total_shots_")
				shoots[weapon] = stat.Value
			}
			if strings.HasPrefix(stat.Name, "total_hits_") {
				weapon := strings.TrimPrefix(stat.Name, "total_hits_")
				hits[weapon] = stat.Value
			}
		}
		for k, v := range shoots {
			accuracy := 0.0
			if _, ok := hits[k]; ok {
				accuracy = (float64(hits[k]) / float64(v)) * 100
			}
			accuracies[k] = accuracy
		}
		return accuracies, nil
	}
	return nil, err
}
