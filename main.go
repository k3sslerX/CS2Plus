package main

import (
	"cs2plus/statsFaceit"
	"cs2plus/statsValve"
	"fmt"
)

func main() {
	user, err := statsValve.FindUser("k3ssler")
	if err == nil {
		faceitUser, err := statsFaceit.GetFaceitPlayer(user.SteamID64)
		if err == nil {
			statistics, err := statsFaceit.GetFaceitLast20MatchesStats(faceitUser)
			fmt.Println(err)
			if err == nil {
				fmt.Println(statistics)
			}
		}

	}
}
