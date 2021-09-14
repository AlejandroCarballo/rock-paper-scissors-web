package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2

)

type Round struct {
	// *** changed Winner to Message
	// Winner         int    `json:"winner"`
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

// *** created three slices of strings, each with exactly three entries
var winMessages = []string{
	"Good job!",
	"Nice work!",
	"You should buy a lottery ticket",
}

var loseMessages = []string{
	"Too bad!",
	"Try again!",
	"This is just not your day.",
}

var drawMessages = []string{
	"Great minds think alike.",
	"Uh oh. Try again.",
	"Nobody wins, but you can try again.",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	// *** generate a random number from 0-2, which we use to pick a random message
	messageInt := rand.Intn(3)
	// *** declare a var to hold the message
	message := ""

	if playerValue == computerValue {
		roundResult = "It's a draw"
		//winner = DRAW
		// *** populate message from drawMessages
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		//winner = PLAYERWINS
		// *** populate message from winMessages
		message = winMessages[messageInt]
	} else {
		roundResult = "Computer wins!"
		//winner = COMPUTERWINS
		// *** populate message from loseMessages
		message = loseMessages[messageInt]
	}

	var result Round
	// *** change to use message instead of Winner
	result.Message = message
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result

}
