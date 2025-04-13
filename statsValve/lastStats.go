package statsValve

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const matchHistoryAPI = "https://api.steampowered.com/ICSGOPlayers_730/GetMatchHistory/v1/"

func GetLastTwoWeeksMapsWinrate(user *SteamUser) (map[string]float64, error) {
	steamAPIKey, err := getApiKey()
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("key", steamAPIKey)
	params.Add("steamid", user.SteamID64)
	params.Add("matches_requested", "100") // Максимум 100 матчей

	resp, err := http.Get(matchHistoryAPI + "?" + params.Encode())
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()

	var data struct {
		Result struct {
			Matches []struct {
				Map    string `json:"map"`
				Result int    `json:"result"`
			} `json:"matches"`
		} `json:"result"`
	}
	json.NewDecoder(resp.Body).Decode(&data)

	mapStats := make(map[string]struct {
		Wins  int
		Total int
	})

	for _, match := range data.Result.Matches {
		stats := mapStats[match.Map]
		stats.Total++
		if match.Result == 1 {
			stats.Wins++
		}
		mapStats[match.Map] = stats
	}
	mapWinRate := make(map[string]float64)
	for k, v := range mapStats {
		mapWinRate[k] = float64(v.Wins) / float64(v.Total)
	}
	return mapWinRate, nil
}
