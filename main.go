package main

import (
	"cs2plus/cs2plus"
	"fmt"
)

func main() {
	user, err := cs2plus.FindUser("k3ssler")
	if err == nil {
		stats, err := cs2plus.GetOverallAccuracy(user)
		if err == nil {
			for k, v := range stats {
				fmt.Printf("%s: %.2f\n", k, v)
			}
		}
	}
}
