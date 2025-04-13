package main

import (
	"cs2plus/cs2plus"
)

func main() {
	user, err := cs2plus.FindUser("k3ssler")
	if err == nil {
		cs2plus.GetOverallStats(user)
	}
	//// Вывод K/D, убийств, смертей и т. д.
	//for _, stat := range data.Playerstats.Stats {
	//	switch stat.Name {
	//	case "total_kills":
	//		fmt.Printf("Убийства: %d\n", stat.Value)
	//	case "total_deaths":
	//		fmt.Printf("Смерти: %d\n", stat.Value)
	//	case "total_mvps":
	//		fmt.Printf("MVP: %d\n", stat.Value)
	//	}
	//}
}
