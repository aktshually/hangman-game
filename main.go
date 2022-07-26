package main

import (
	"fmt"
	"hangman/utils"
	"strings"
)

func main() {
	fmt.Println("Get ready! We're preparing everything.")
	game := utils.Game{}
	game.New()

	fmt.Println("Your game has been loaded! Let's begin!")

	for true {
		game.ShowStats()

		letter := game.InputLetter()
		if game.Word.ContainsLetter(letter) && !game.AlreadyTried.ContainsLetter(letter) {
			game.AlreadyTried.AddLetter(letter)
		} else if game.AlreadyTried.ContainsLetter(letter) {
			fmt.Println("You've already tried to put this letter! Try again.")
			continue
		} else {
			game.ChancesLeft.DecreaseByOne()
		}

		if game.IsCompleted() {
			break
		}
	}

	if !strings.Contains(game.Word.Censor(game.AlreadyTried.ContainsLetter), "_") {
		fmt.Println("You won the game! Congratulations!")
	} else if game.ChancesLeft <= 0 {
		fmt.Println("You have no more chances. Good luck in the next time!")
	}
}
