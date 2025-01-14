# What is this?
Some years ago, I was a beginner in the Go programming language and had the "brilliant" idea of creating a hangman game that could be played from your terminal.
I, some years later, still think it's a cool idea, so this README aims to provide instructions for you to run this project on your own machine.

# Requisites
- [Go](https://go.dev/)

# How to run
I haven't tested this for newer Go versions (even though I don't think any of the new versions' changes would break the code), so if you have any issues running this code, **I highly recommend you downgrade your Go version to 1.18**, which is the version I used to develop this minigame.
To run the game, simply clone this repository:
```bash
git clone https://github.com/aktshually/hangman-game.git
```
Go to the cloned directory:
```bash
cd hangman-game
```
And run the code as you would do for any Go program:
```bash
go run ./main.go
```
  The game depends on [this API](https://random-word-api.herokuapp.com/word), so it won't work if this is shut down.
