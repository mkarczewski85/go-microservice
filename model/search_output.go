package model

type WordSearchOutput struct {
	word          string `json:"word"`
	scrabbleScore uint16 `json:"scrabbleScore"`
}
