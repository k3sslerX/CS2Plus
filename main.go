package main

import (
	"cs2plus/statsFaceit"
	"cs2plus/statsValve"
	"fmt"
	"strings"
)

func PrintFaceitStats(stats *statsFaceit.FaceitStats) {
	// Создаем красивую таблицу с выравниванием
	header := "FACEIT STATISTICS"
	separator := strings.Repeat("=", 40)

	fmt.Printf("\n%s\n%s\n", header, separator)

	// Форматируем float значения с 2 знаками после запятой
	fmt.Printf("%-25s %15.2f%%\n", "Winrate:", stats.Winrate)
	fmt.Printf("%-25s %15d\n", "Total Kills:", stats.TotalKills)
	fmt.Printf("%-25s %15d\n", "Total Headshots:", stats.TotalHeadshots)
	fmt.Printf("%-25s %15d\n", "Total Assists:", stats.TotalAssists)
	fmt.Printf("%-25s %15d\n", "Total Deaths:", stats.TotalDeaths)
	fmt.Printf("%-25s %15d\n", "Total MVPs:", stats.TotalMVPs)
	fmt.Printf("%-25s %15d\n", "Avg Kills per Match:", stats.AvgKills)
	fmt.Printf("%-25s %15d\n", "Avg Headshots per Match:", stats.AvgHeadshots)
	fmt.Printf("%-25s %15d\n", "Avg Assists per Match:", stats.AvgAssists)
	fmt.Printf("%-25s %15d\n", "Avg MVPs per Match:", stats.AvgMVPs)
	fmt.Printf("%-25s %15.2f\n", "Avg K/D Ratio:", stats.AvgKD)
	fmt.Printf("%-25s %15.2f\n", "Avg K/R Ratio:", stats.AvgKR)
	fmt.Printf("%-25s %15.2f\n", "Average ADR:", stats.ADR)
	fmt.Println(separator)
}

func main() {
	user, err := statsValve.FindUser("k3ssler")
	if err == nil {
		faceitUser, err := statsFaceit.GetFaceitPlayer(user.SteamID64)
		if err == nil {
			statistics, err := statsFaceit.GetFaceitLast20MatchesStats(faceitUser)
			if err == nil {
				PrintFaceitStats(statistics)
			} else {
				fmt.Println(err)
			}
		}

	}
}
