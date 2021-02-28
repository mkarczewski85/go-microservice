package model

type WordsSearchInput struct {
	Letters         string `json:"letters" binding:"required"`
	AuxiliaryLetter string `json:"auxiliaryLetter"`
	Mode            uint8  `json:"mode" binding:"required"`
}
