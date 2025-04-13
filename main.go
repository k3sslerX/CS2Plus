package main

import (
	"cs2plus/cs2plus"
)

func main() {
	user, err := cs2plus.FindUser("k3ssler")
	if err == nil {
		cs2plus.GetOverallStats(user)
	}
}
