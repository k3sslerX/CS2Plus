package statsValve

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const resolveVanityAPI = "https://api.steampowered.com/ISteamUser/ResolveVanityURL/v1/"

type SteamUser struct {
	SteamID64 uint64 `json:"steamid64"`
	CustomURL string `json:"custom_url"`
}

type VanityResponse struct {
	Response struct {
		SteamID string `json:"steamid"`
		Success int    `json:"success"`
	} `json:"response"`
}

func FindUser(link string) (*SteamUser, error) {
	steamAPIKey, err := getApiKey()
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("key", steamAPIKey)
	params.Add("vanityurl", link)

	resp, err := http.Get(resolveVanityAPI + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data VanityResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if data.Response.Success == 1 {
		steamID, err := strconv.ParseUint(data.Response.SteamID, 10, 64)
		if err != nil {
			return nil, err
		}
		user := SteamUser{
			SteamID64: steamID,
			CustomURL: link,
		}
		return &user, nil
	} else {
		return nil, fmt.Errorf("error in response: %s", string(body))
	}
}
