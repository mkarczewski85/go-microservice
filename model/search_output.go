package model

type WordSearchOutput struct {
	Word          string `json:"word"`
	ScrabbleScore uint16 `json:"scrabbleScore"`
}
