package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type RandomWordResponse []string

func GetRandomWord() string {
	response, err := http.Get("https://random-word-api.herokuapp.com/word")
	if err != nil {
		panic("the word couldn't be fetched")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic("could not parse the response body")
	}

	randomWord := RandomWordResponse{}
	json.Unmarshal(body, &randomWord)

	return randomWord[0]
}
