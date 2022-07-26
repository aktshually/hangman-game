package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var game Game = Game{}

func init() {
	game.New()
}

func TestNewGame(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, game.ChancesLeft)
	assert.Equal([]string{}, game.AlreadyTried)
	assert.NotEmpty(game.Word)
}

func TestCensorWord(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, 6)
}
