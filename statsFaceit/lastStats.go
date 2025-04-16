package statsFaceit

// RESULT 0 - L / 1 - W

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"strconv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func getApiKey() (string, error) {
	faceitAPIKey, flag := os.LookupEnv("FACEIT_API")
	if !flag {
		fmt.Println("FACEIT_API not set")
		return "", errors.New("FACEIT_API not set")
	}
	return faceitAPIKey, nil
}

type FaceitStats struct {
	Winrate          float64
	TotalKills       int
	TotalHeadshots   int
	TotalAssists     int
	TotalDeaths      int
	TotalMVPs        int
	AvgKills         int
	AvgHeadshots     int
	HeadshotsPercent float64
	AvgAssists       int
	AvgMVPs          int
	AvgKD            float64
	AvgKR            float64
	ADR              float64
}

type MatchStats struct {
	Kills     string `json:"Kills"`
	Headshots string `json:"Headshots"`
	Assists   string `json:"Assists"`
	Deaths    string `json:"Deaths"`
	MVPs      string `json:"MVPs"`
	ADR       string `json:"ADR"`
	KRRatio   string `json:"K/R Ratio"`
	Result    string `json:"Result"`
}

type Item struct {
	Stats MatchStats `json:"stats"`
}

type Response struct {
	Items []Item `json:"items"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

func GetFaceitLast20MatchesStats(user *FaceitUser) (*FaceitStats, error) {
	url := fmt.Sprintf("https://open.faceit.com/data/v4/players/%s/games/cs2/stats", user.FaceitID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	faceitAPIKey, err := getApiKey()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+faceitAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	stats := FaceitStats{}
	adrs := 0.0
	krs := 0.0
	wins := 0
	for _, item := range response.Items {
		kills, _ := strconv.Atoi(item.Stats.Kills)
		headshots, _ := strconv.Atoi(item.Stats.Headshots)
		assists, _ := strconv.Atoi(item.Stats.Assists)
		deaths, _ := strconv.Atoi(item.Stats.Deaths)
		mvps, _ := strconv.Atoi(item.Stats.MVPs)
		stats.TotalKills += kills
		stats.TotalHeadshots += headshots
		stats.TotalAssists += assists
		stats.TotalDeaths += deaths
		stats.TotalMVPs += mvps

		adr, _ := strconv.ParseFloat(item.Stats.ADR, 64)
		adrs += adr
		kr, _ := strconv.ParseFloat(item.Stats.KRRatio, 64)
		krs += kr
		win, _ := strconv.Atoi(item.Stats.Result)
		wins += win
	}
	stats.AvgKills = stats.TotalKills / len(response.Items)
	stats.AvgHeadshots = stats.TotalHeadshots / len(response.Items)
	stats.AvgAssists = stats.TotalAssists / len(response.Items)
	stats.AvgMVPs = stats.TotalMVPs / len(response.Items)
	stats.AvgKD = float64(stats.TotalKills) / float64(stats.TotalDeaths)
	stats.AvgKR = krs / float64(len(response.Items))
	stats.ADR = adrs / float64(len(response.Items))
	stats.HeadshotsPercent = (float64(stats.TotalHeadshots) / float64(stats.TotalKills)) * 100
	stats.Winrate = (float64(wins) / float64(len(response.Items))) * 100.0
	return &stats, nil
}
