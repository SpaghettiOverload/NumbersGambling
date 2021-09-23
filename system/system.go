package system

import (
	"fmt"
	F "main/functions"
	"main/player"
	"os"
	"strconv"
	"strings"
)

const (
	err = "ERROR"
	hint = "HINT"
	wrongGuess = "WRONG GUESS"
	correctGuess = "CORRECT"
	wrongGuessTax = "$50"
	winingStrikeFailureTax = "50% from your balance"
	successfulSale = "Successful purchase"
	blueColor = "\033[34m"
	greenColor = "\033[32m"
	redColor = "\033[31m"
	resetColor = "\033[0m"
)

type System struct {
	Range int
	Round int
	continueRound bool
	hintGiven bool
	winningStrike bool
	playerCommand string
	CurrentNum int
	nextNum   int
	isInt     bool
	intResult int
	ValidCommands []string
	Player player.Player
	Shop map[string]int
}

// negativeMessage returns a message colored in red.
func (s *System) negativeMessage(msg string) string {
	return redColor + msg + resetColor
}

// positiveMessage returns a message colored in green.
func (s *System) positiveMessage(msg string) string {
	return greenColor + msg + resetColor
}

// roundOn keeps a loop of taking/evaluating user's input going on for the current round.
func (s *System) roundOn() {
	s.takingInput()
	s.evaluateCommand()
}

// GameOn is where the logic is assembled and keeps the rounds rotating.
func (s *System) GameOn() {
	for {
		fmt.Println(s.gameRound() + " / " + s.Player.State())
		s.ValidCommands = []string{"H", "U", "D", "J", "R", "B"}
		s.continueRound = true
		s.hintGiven = false
		F.WaitLonger()
		s.nextNum = F.NewRandomNum(s.Range)

		// Handling the case where the new number might be equal to the old.
		for s.nextNum == s.CurrentNum {
			s.nextNum = F.SpareRandomNum(s.Range)
		}

		for s.continueRound == true {
			s.roundOn()
		}
		s.CurrentNum = s.nextNum
	}
}

// chargedFor deducts the needed amount for the purchased item from the player's balance and prints confirmation or error.
func (s *System) chargedFor(item string) bool {
	if s.Player.Balance - s.Shop[item] >= 0 {
		s.Player.Balance -= s.Shop[item]
		fmt.Print(s.positiveMessage(successfulSale) + " of " + strings.ToUpper(item) + " / " + s.Player.State())
		F.WaitLonger()
		return true
	}
	fmt.Print(s.negativeMessage(err) + ": Not enough balance!")
	F.WaitLonger()
	return false
}

// gameRound autoincrement round count and returns it as current round.
func (s *System) gameRound() string {
	s.Round += 1
	return blueColor + "ROUND " + strconv.Itoa(s.Round) + resetColor + "\nBase number: " + strconv.Itoa(s.CurrentNum)
}

// takingInput takes care to obtain valid user's input and sets it as system's PlayerCommand for the current round.
func (s *System) takingInput() {  //
	s.playerCommand = s.takeInput()
	for {
		if s.validInput(s.playerCommand) {
			break
		}
		fmt.Print("Ooops, invalid command! Try again!")
		F.WaitLonger()
		F.CleanLine()
		s.playerCommand = s.takeInput()
	}
}

// takeInput asks for input and provides basic "no empty input" validation.
// It returns whatever else in upper case for further processing.
func (s *System) takeInput() string {
	var n string
	for {
		fmt.Print("Enter your command/guess: ")
		_, err := fmt.Scanln(&n)
		if err != nil {
			fmt.Print("Ooops, value needed!")
			F.WaitLonger()
			F.CleanLine()
			continue
		}
		break
	}
	return strings.ToUpper(n)
}

// validInput validates whether the input from TakeInput is a whole number or a valid string command.
func (s *System) validInput(userInput string) bool {
	value, err := strconv.Atoi(userInput)
	if err == nil {
		s.intResult = value
		s.isInt = true
	}
	if !s.isInt {
		if !s.validStringCommand() {
			return false
		}
	}
	return true
}

// validStringCommand evaluates user's non-numerical input against the system's ValidCommands array.
func (s *System) validStringCommand() bool {
	for _, v := range s.ValidCommands {
		if v == s.playerCommand {
			return true
		}
	}
	return false
}

// evaluateCommand applies logic according to the validated user's input.
func (s *System) evaluateCommand()  {
	switch {
	case s.isInt == true:
		s.isInt = false
		s.jackpotGuess()
	case s.playerCommand == "H":
		if s.hintGiven {
			fmt.Print(s.negativeMessage(err) + ": Only 1 hint per round is allowed!")
			F.WaitLonger()
		} else {
			fmt.Print(s.hintMessage())
			s.Player.Hints -= 1
			s.hintGiven = true
		}
	case s.playerCommand == "J":
		if s.chargedFor("Jackass credit") {
			s.Player.JackassCredit += 1
		}
	case s.playerCommand == "B":
		fmt.Print(s.Player.State())
	case s.playerCommand == "R":
		if s.Range > 10 {
			if s.chargedFor("Range reduction") {
				s.continueRound = false
				s.Range -= 10
				}
			} else {
				fmt.Print(s.negativeMessage(err) + ": Range is already at minimum!")
			}
	case s.playerCommand == "C" && s.winningStrike == true:
		s.winningStrike = false
		s.continueRound = false
		s.Player.Balance += s.Player.TempBalance
		fmt.Print("CASHING OUT: $")
		fmt.Print(s.Player.TempBalance)
	case s.playerCommand == "U":
		if s.nextNum > s.CurrentNum {
			if !s.winningStrike {
				s.winingStrike()
			}
		} else {
			fmt.Println(s.wrongGuess())
			s.continueRound = false
		}
	case s.playerCommand == "D":
		if s.nextNum < s.CurrentNum {
			if !s.winningStrike {
				s.winingStrike()
			}
		} else {
			fmt.Println(s.wrongGuess())
			s.continueRound = false
		}
	}
	F.CleanLine()
}

// hintMessage returns a message whether the newly generated number went higher or lower compared to the previous one.
// It also deducts player's hint credits (if such) or returns an error.
func (s *System) hintMessage() string {
	if s.Player.Hints > 0 {
		msg := s.positiveMessage(hint) + ": The new number is"
		h := " lower"
		if s.nextNum > s.CurrentNum {
			h = " higher"
		}
		return msg + h + " than " + strconv.Itoa(s.CurrentNum) + "\nHINT credits left: " + strconv.Itoa(s.Player.Hints)
	}
	return s.negativeMessage(err) + ": You have no HINT credits"
}

// jackpotGuess handles WIN scenario (if such) by ending the program or returns an error message if the guess is wrong.
// It also deducts player's JACKASS credits (if such) or returns an error.
func (s *System) jackpotGuess() {
	if s.Player.JackassCredit > 0 {
		s.Player.JackassCredit -= 1
		if s.intResult == s.nextNum {
			fmt.Println(s.Player.Win())
			os.Exit(0)
		}
		fmt.Print(s.wrongGuess() + "\nJACKASS credits left: " + strconv.Itoa(s.Player.JackassCredit) + " / " + s.Player.State())
	}
	fmt.Print(s.negativeMessage(err) + ": Your JACKASS credits are depleted")
	F.WaitLonger()
}

// lose terminates the program in a case of a loss.
func (s *System) lose() {
	fmt.Println(s.Player.Lose())
	os.Exit(1)
}

// wrongGuess handles cases of wrong guesses and charge respective amount to player's balance.
// It returns confirmation message of the action taken.
// It also can end the program in a losing scenario if balance become < 0 by calling system's lose function.
func (s *System) wrongGuess() string {
	if s.winningStrike {
		s.Player.Balance /= 2
	}
	s.Player.Balance -= 50
	if s.Player.Balance < 0 {
		s.lose()
	}
	if s.winningStrike {
		s.winningStrike = false
		return s.negativeMessage(wrongGuess+"!") + " You've been fined " + winingStrikeFailureTax + " + " + wrongGuessTax
	}
	return s.negativeMessage(wrongGuess+"!") + " You've been fined " + wrongGuessTax
}

// correctGuess returns a green colored message.
func (s *System) correctGuess() string {
	return s.positiveMessage(correctGuess+"!")
}

// winingStrike handles Wining Strike State encapsulated from the main logic.
func (s *System) winingStrike() {
	s.ValidCommands = []string{"U", "D", "C"}
	s.Player.TempBalance = 200
	s.winningStrike = true
	fmt.Println(s.correctGuess() + " The number was " + strconv.Itoa(s.nextNum))
	fmt.Print("Entering Wining Strike State!")
	F.CleanLine()
	stage := 1
	for s.winningStrike == true {
		fmt.Println("WS Stage - " + strconv.Itoa(stage) + " / PRIZE - $" + strconv.Itoa(s.Player.TempBalance))
		s.rotateNums()
		s.takingInput()
		s.evaluateCommand()
		if s.winningStrike {
			fmt.Println(s.correctGuess() + " The number was " + strconv.Itoa(s.nextNum))
			stage += 1
			s.Player.TempBalance *= 4
		}
	}
}

// rotateNums makes last unknown number known while generates new unknown number.
// Serves winingStrike functionality.
func (s *System) rotateNums() {
	s.CurrentNum = s.nextNum
	s.nextNum = F.NewRandomNum(s.Range)
}
