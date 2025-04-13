package statsFaceit

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type FaceitUser struct {
	SteamID64 string `json:"steamid64"`
	CustomURL string `json:"custom_url"`
	FaceitID  string `json:"faceit_id"`
}

func GetFaceitPlayer(steamID string) (*FaceitUser, error) {
	baseURL := "https://open.faceit.com/data/v4/players"

	params := url.Values{}
	params.Add("game", "cs2")
	params.Add("game_player_id", steamID)

	req, err := http.NewRequest("GET", baseURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	faceitAPIKey, err := getApiKey()
	req.Header.Set("Authorization", "Bearer "+faceitAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		PlayerID string `json:"player_id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	faceitUser := FaceitUser{
		SteamID64: steamID,
		FaceitID:  result.PlayerID,
	}
	return &faceitUser, nil
}
