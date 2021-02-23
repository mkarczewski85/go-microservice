package model

type WordsSearchInput struct {
	Letters         []rune `json:"letters" binding:"required"`
	AuxiliaryLetter rune   `json:"auxiliaryLetter" binding:"required"`
	Option          uint8  `json:"option" binding:"required"`
}
