package utils

import (
	"fmt"
	"hangman/api"
	"strings"
	"unicode"
)

type Game struct {
	ChancesLeft  Chances
	Word         Word
	AlreadyTried AlreadyTried
}

func (game *Game) New() {
	game.ChancesLeft = Chances(6)
	game.Word = Word(api.GetRandomWord())
	game.AlreadyTried = []string{}
}

func (game *Game) ShowStats() {
	censoredWord := game.Word.Censor(game.AlreadyTried.ContainsLetter)

	fmt.Println()
	fmt.Println(censoredWord)
	fmt.Printf("Chances left: %d\n", game.ChancesLeft)
	fmt.Printf("Letters you tried: %s", strings.Join(game.AlreadyTried, " "))
	fmt.Println()
}

func (game *Game) InputLetter() string {
	var letter string

	for letter == "" {
		var input string
		fmt.Print("Enter a letter: ")

		fmt.Scanln(&input)
		if len(input) > 1 {
			fmt.Println("Oops, you have to type only one letter!")
			continue
		}

		reader := strings.NewReader(input)
		r, _, _ := reader.ReadRune()
		if !unicode.IsLetter(r) {
			fmt.Println("You can't use symbols/numbers!")
			continue
		}

		letter = input
	}

	return letter
}

func (game *Game) IsCompleted() bool {
	return !strings.Contains(game.Word.Censor(game.AlreadyTried.ContainsLetter), "_") || game.ChancesLeft <= 0
}
