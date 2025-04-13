package main

import (
	"cs2plus/statsValve"
	"fmt"
)

func main() {
	user, err := statsValve.FindUser("k3ssler")
	if err == nil {
		statistics, err := statsValve.GetOverallAccuracy(user)
		fmt.Println(err)
		if err == nil {
			for k, v := range statistics {
				fmt.Printf("%s: %.2f\n", k, v)
			}
		}
	}
}
