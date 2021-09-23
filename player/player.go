package player

import (
	"strconv"
)

const (
	loseMessage = "YOU HAVE LOST THE GAME!\n"
	winMessage = "CONGRATULATIONS!\nYou have won the game!\n"
	greenColor = "\033[32m"
	redColor = "\033[31m"
	resetColor = "\033[0m"
)

type Player struct {
	Balance  int
	TempBalance int
	JackassCredit int
	Hints int
}

func (p *Player) Win() string {
	p.Balance = 1000000
	return greenColor + winMessage + resetColor + p.State() + "\nNow get yourself a beer and CHEERS!"
}

func (p *Player) Lose() string {
	return redColor + loseMessage + resetColor + p.State()
}

func (p *Player) State() string {
	return "Your balance: $" + strconv.Itoa(p.Balance)
}
