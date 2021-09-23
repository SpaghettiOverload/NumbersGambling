package main

import (
	"fmt"
	F "main/functions"
	"main/player"
	"main/system"
)

func main() {
	fmt.Print("Creating a game session")
	F.Loading()  // Let's fake some heavy logic here :)
	randomNum := F.SpareRandomNum(100)

	sys := system.System{
		Range: 100,
		Round: 0,
		CurrentNum: randomNum,
		Player: player.Player{
			Balance: 100,
			JackassCredit: 1,
			Hints: 3},
		Shop: map[string]int{
			"Jackass credit": 2500,
			"Range reduction": 250000,
		},
	}
	fmt.Print("OK, Let's roll!")
	F.Wait()
	F.CleanLine()
	sys.GameOn()
}
