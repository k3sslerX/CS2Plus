package statsFaceit

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
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

type MapStats struct {
	MapName string
	Wins    int
	Losses  int
}

func GetFaceitLast20MatchesStats(user *FaceitUser) ([]MapStats, error) {
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
	fmt.Println(string(body))

	return nil, nil
}
